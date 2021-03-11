package vars

import "os"

func VariableString(name string, value string) {
	_ = os.Setenv(name, value[1:len(value)-1])
}

func VariableInt(name string, value int) {

}

func VariableFloat(name string, value float64) {

}
