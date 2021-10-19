package configuration

import (
	"errors"
	"flag"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func (c *ConfigFromYaml) GetOptions() (r RequestConfig, err error) {
	var file []byte
	file, err = os.ReadFile(Path)

	if err != nil{
		err = errors.New("ошибка открытия конфигурационного файла")
		return r, err
	}

	if err = yaml.Unmarshal(file, &c); err != nil{
		err = errors.New("ошибка чтения конфигурационного файла")
		return r, err
	}

	r.TownName = c.Name
	r.UnitSystem = c.Units
	return r, nil
}


func (y *ConfigFromFlags) GetOptions() (r RequestConfig, err error)  {
	t := flag.String(TownFlagName, DefaultTownFlag, "choose town")
	u := flag.String(UnitsFlagName, DefaultUnitsFlag, "standard, metric and imperial")
	flag.Parse()
	y.T, y.U = *t, *u

	var falseOptions bool
	for _, v := range []string{"metric", "standard", "imperial"} {
		if y.U == v {
			falseOptions = true
			break
		}
	}
	if falseOptions != true{
		err = errors.New("некорректное значение, используейте -h")
		return r, err
	}

	r.TownName = y.T
	r.UnitSystem = y.U
	return r, nil
}

func (r *RequestConfig) ConfigSelector(y *RequestConfig, f *RequestConfig) () {
	log.Printf("YAML config is: %v\n", *y)
	log.Printf("FLAG config is: %v\n", *f)

	//должен сравнить два конфига, если флаги не по умолчанию, то выбрать флаги. В остальных случаях выбрать yaml
	if f.TownName == DefaultTownFlag && f.UnitSystem == DefaultUnitsFlag {
		r.TownName = y.TownName
		r.UnitSystem = y.UnitSystem
	} else {
		r.TownName = f.TownName
		r.UnitSystem = f.UnitSystem
	}
}

type OptionsGetter interface {
	GetOptions() (*RequestConfig, error)
}