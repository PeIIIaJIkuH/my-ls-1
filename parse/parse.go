package parse

import (
	"errors"
	"fmt"
	"my-ls-1/models"
	"os"
	"os/user"
	"syscall"
)

func Args(args []string) ([]string, *models.Flags, error) {
	var flags models.Flags
	var filenames []string
	for _, arg := range args {
		if arg[0] == '-' {
			if len(arg) == 1 {
				return nil, nil, errors.New(fmt.Sprintf(models.IncorrectFlags, ""))
			}
			for _, c := range arg[1:] {
				switch c {
				case 'a':
					flags.All = true
				case 'r':
					flags.Reverse = true
				case 'R':
					flags.Recursive = true
				case 'l':
					flags.Long = true
				case 't':
					flags.SortByTime = true
				default:
					return nil, nil, errors.New(fmt.Sprintf(models.IncorrectFlags, string(c)))
				}
			}
		} else {
			filenames = append(filenames, arg)
		}
	}
	if len(filenames) == 0 {
		filenames = append(filenames, ".")
	}
	return filenames, &flags, nil
}

func Filenames(filenames []string) error {
	for _, filename := range filenames {
		if _, err := os.Stat(filename); err != nil {
			return errors.New(fmt.Sprintf(models.NotExist, filename))
		}
	}
	return nil
}

func New(stat os.FileInfo) (*models.Entity, error) {
	var entity models.Entity
	entity.Name = stat.Name()
	entity.Permissions = stat.Mode().String()
	entity.ModTime = stat.ModTime()

	//if time.Now().Year() == entity.ModTime.Year() {
	//	fmt.Printf("%s %2d %02d:%02d\n", strings.ToLower(entity.ModTime.Format("Jan")), entity.ModTime.Day(), entity.ModTime.Hour(), entity.ModTime.Minute())
	//} else {
	//	fmt.Printf("%s %2d %5d\n", strings.ToLower(entity.ModTime.Format("Jan")), entity.ModTime.Day(), entity.ModTime.Year())
	//}

	if sys := stat.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			entity.HardLinks = stat.Nlink
			userInfo, err := user.LookupId(fmt.Sprintf("%d", stat.Uid))
			if err != nil {
				return nil, err
			}
			entity.UserOwner = userInfo.Username
			groupInfo, err := user.LookupGroupId(fmt.Sprintf("%d", stat.Uid))
			if err != nil {
				return nil, err
			}
			entity.GroupOwner = groupInfo.Name
		}
	}
	return &entity, nil
}

func Entities(fileNames []string, parentName string, flags *models.Flags) ([]models.Entity, error) {
	var entities []models.Entity
	for _, filename := range fileNames {
		totalName := filename
		if parentName != "" {
			totalName = parentName + "/" + filename
		}
		info, err := os.Stat(totalName)
		if err != nil {
			return nil, err
		}
		entity, err := New(info)
		if err != nil {
			return nil, err
		}
		if info.IsDir() {
			files, err := os.ReadDir(totalName)
			if err != nil {
				return nil, err
			}
			for _, file := range files {
				if !flags.All && file.Name()[0] == models.DotCharacter {
					continue
				}
				innerInfo, err := file.Info()
				if err != nil {
					return nil, err
				}
				innerEntity, err := New(innerInfo)
				if err != nil {
					return nil, err
				}
				if flags.Recursive && file.IsDir() {
					innerTotalName := totalName + "/" + file.Name()
					innerFiles, err := os.ReadDir(innerTotalName)
					if err != nil {
						return nil, err
					}
					var innerFilenames []string
					for _, innerFile := range innerFiles {
						innerFilenames = append(innerFilenames, innerFile.Name())
					}
					newEntities, err := Entities(innerFilenames, innerTotalName, flags)
					if err != nil {
						return nil, err
					}
					innerEntity.Children = append(innerEntity.Children, newEntities...)
				}
				entity.Children = append(entity.Children, *innerEntity)
			}
		}
		entities = append(entities, *entity)
	}
	return entities, nil
}
