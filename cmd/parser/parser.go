package parser

//file of guide that I saw on a tutorial. If you don't include it in the init func
// from root it won't do anything
import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
)

var (
	ErrTokenEmpty    = errors.New("token is empty")
	ErrTokenNotValid = errors.New("token must have 3 parts")
)
var ParserCmd = &cobra.Command{
	Use:   "parse",
	Short: "parses a jwt without validation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil //por el momento
	},
}

func parseToken(token string) error {
	if token == "" {
		return ErrTokenEmpty
	}
	tokenParts := strings.Split(token, ".")

	if len(tokenParts) != 3 {
		return ErrTokenNotValid
	}
	//here you should do more stuff but I understood the fluency of cobra
	return nil
}
