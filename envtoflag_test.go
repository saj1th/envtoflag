package envtoflag

import (
	"flag"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	os.Setenv("ENVTOFLAG_NAME", "Jattir")
	os.Setenv("ENVTOFLAG_GROUP_NAME", "Jattir Collective")

	var flagName = flag.String("name", "", "")
	var flagGroup = flag.String("group.name", "", "")

	if err := parse("envtoflag"); err != nil {
		t.Error(err)
	}

	if *flagName != "Jattir" {
		t.Fail()
	}

	if *flagGroup != "Jattir Collective" {
		t.Fail()
	}

}
