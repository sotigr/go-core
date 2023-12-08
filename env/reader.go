package env

import (
	"log"
	"os"
	"strconv"
)

type Var struct {
	Name  string
	Value *string
}

func (v Var) toString() (string, error) {
	return *v.Value, nil
}

func (v Var) toFloat64() (float64, error) {
	c, err := strconv.ParseFloat(*v.Value, 64)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func (v Var) toInt() (int, error) {
	c, err := strconv.Atoi(*v.Value)
	if err != nil {
		return 0, err
	}
	return c, nil
}

func (v Var) toBool() (bool, error) {
	c, err := strconv.ParseBool(*v.Value)
	if err != nil {
		return false, err
	}
	return c, nil
}

func readEnv(name string, value string) *Var {
	v := os.Getenv(name)

	if len(v) == 0 {
		return &Var{Name: name, Value: &value}
	}
	return &Var{Name: name, Value: &v}
}

type EnvReader struct{}

func NewEnvReader() *EnvReader {
	return &EnvReader{}
}

func (e EnvReader) StringVar(p *string, name string, value string) error {
	v, err := readEnv(name, value).toString()
	if err != nil {
		return err
	} else {
		*p = v
		return nil
	}
}

func (e EnvReader) String(name string, value string) (string, error) {
	v, err := readEnv(name, value).toString()
	if err != nil {
		return "", err
	} else {
		return v, nil
	}
}

func (e EnvReader) IntVar(p *int, name string, value int) error {
	v, err := readEnv(name, strconv.Itoa(value)).toInt()
	if err != nil {
		return err
	} else {
		*p = v
		return nil
	}
}

func (e EnvReader) Int(name string, value int) (int, error) {
	v, err := readEnv(name, strconv.Itoa(value)).toInt()
	if err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

func (e EnvReader) Float64Var(p *float64, name string, value float64) error {
	fstr := strconv.FormatFloat(value, 'f', 6, 64)
	v, err := readEnv(name, fstr).toFloat64()
	if err != nil {
		return err
	} else {
		*p = v
		return nil
	}
}

func (e EnvReader) Float64(name string, value float64) (float64, error) {
	fstr := strconv.FormatFloat(value, 'f', 6, 64)
	v, err := readEnv(name, fstr).toFloat64()
	if err != nil {
		return 0, err
	} else {
		return v, nil
	}
}

func (e EnvReader) BoolVar(p *bool, name string, value bool) error {
	v, err := readEnv(name, strconv.FormatBool(value)).toBool()
	if err != nil {
		return err
	} else {
		*p = v
		return nil
	}
}

func (e EnvReader) Bool(name string, value bool) (bool, error) {
	v, err := readEnv(name, strconv.FormatBool(value)).toBool()
	if err != nil {
		return false, err
	} else {
		return v, nil
	}
}

type EnvReaderOrExit struct{}

func NewEnvReaderOrExit() *EnvReaderOrExit {
	return &EnvReaderOrExit{}
}

func (e EnvReaderOrExit) StringVar(p *string, name string, value string, usage string) {
	v, err := readEnv(name, value).toString()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
	} else {
		*p = v
	}
}

func (e EnvReaderOrExit) String(name string, value string, usage string) string {
	v, err := readEnv(name, value).toString()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
		return ""
	} else {
		return v
	}
}

func (e EnvReaderOrExit) IntVar(p *int, name string, value int, usage string) {
	v, err := readEnv(name, strconv.Itoa(value)).toInt()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
	} else {
		*p = v
	}
}

func (e EnvReaderOrExit) Int(name string, value int, usage string) int {
	v, err := readEnv(name, strconv.Itoa(value)).toInt()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
		return 0
	} else {
		return v
	}
}

func (e EnvReaderOrExit) Float64Var(p *float64, name string, value float64, usage string) {
	fstr := strconv.FormatFloat(value, 'f', 6, 64)
	v, err := readEnv(name, fstr).toFloat64()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
	} else {
		*p = v
	}
}

func (e EnvReaderOrExit) Float64(name string, value float64, usage string) float64 {
	fstr := strconv.FormatFloat(value, 'f', 6, 64)
	v, err := readEnv(name, fstr).toFloat64()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
		return 0
	} else {
		return v
	}
}

func (e EnvReaderOrExit) BoolVar(p *bool, name string, value bool, usage string) {
	v, err := readEnv(name, strconv.FormatBool(value)).toBool()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
	} else {
		*p = v
	}
}

func (e EnvReaderOrExit) Bool(name string, value bool, usage string) bool {
	v, err := readEnv(name, strconv.FormatBool(value)).toBool()
	if err != nil {
		log.Fatal("Fatal error: invalid usage of " + name + ". " + usage)
		return false
	} else {
		return v
	}
}
