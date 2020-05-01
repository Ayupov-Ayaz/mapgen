package cmd

import (
	"errors"
	"log"

	"github.com/ayupov-ayaz/mapgen/analysis"
	"github.com/ayupov-ayaz/mapgen/services"

	"github.com/spf13/cobra"
)

const (
	cmdMapBySlice = "map_by_slice"
)

var (
	runCmd = &cobra.Command{
		Use:   "run",
		Short: "generate map to switch case",
		RunE:  run,
	}
	runFlags = struct {
		Package       string
		Path          string
		StructName    string
		SearchMapType string
		CountType     string
		Command       string
	}{}
)

func init() {
	runCmd.PersistentFlags().StringVarP(&runFlags.Package, "package", "p",
		"", "package name for generated struct")

	runCmd.PersistentFlags().StringVarP(&runFlags.Path, "filepath", "f",
		"", "path to file where i can find map")

	runCmd.PersistentFlags().StringVarP(&runFlags.StructName, "name", "n",
		"", "for new generated struct name")

	runCmd.PersistentFlags().StringVarP(&runFlags.CountType, "arg_type", "a",
		"uint8", "count type (uint8, uint16, uint32, int, int32, int64)")

	runCmd.PersistentFlags().StringVarP(&runFlags.SearchMapType, "type", "t",
		"", "map type struct; Example: t=MapPayout for type MapPayout = map[string][]int")

	runCmd.PersistentFlags().StringVarP(&runFlags.Command, "command", "c", cmdMapBySlice,
		"generate command")
}

func validateFlags() error {
	if len(runFlags.Package) == 0 {
		return errors.New("'package' flag doesn't be empty")
	}

	if len(runFlags.SearchMapType) == 0 {
		return errors.New("'type' flag doesn't be empty")
	}

	if len(runFlags.StructName) == 0 {
		return errors.New("'name' flag doesn't be empty")
	}

	if len(runFlags.Path) == 0 {
		return errors.New("'path' flag doesn't be empty")
	}

	return nil
}

func checkCountType() error {
	switch runFlags.CountType {
	case "uint8", "uint16", "uint32", "uint64", "int", "int32", "int64":
		return nil
	}
	return errors.New("invalid count type")
}

func run(cmd *cobra.Command, args []string) error {
	err := validateFlags()
	if err != nil {
		return err
	}

	err = checkCountType()
	if err != nil {
		return err
	}

	recorder := services.NewRecorder()
	switch runFlags.Command {
	case cmdMapBySlice:
		mp := analysis.NewMapParams(
			runFlags.Package,
			runFlags.Path,
			runFlags.SearchMapType,
			runFlags.StructName,
			runFlags.CountType,
		)

		err = analysis.GenerateMapByString(recorder, mp)
	default:
		err = errors.New("command not supported")
	}

	return err
}

func Execute() {
	if err := runCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
