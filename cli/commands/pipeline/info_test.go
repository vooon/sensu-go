package pipeline

import (
	"fmt"
	"testing"

	corev2 "github.com/sensu/sensu-go/api/core/v2"
	client "github.com/sensu/sensu-go/cli/client/testing"
	test "github.com/sensu/sensu-go/cli/commands/testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfoCommand(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Config.(*client.MockConfig).On("Format").Return("json")
	cmd := InfoCommand(cli)

	assert.NotNil(t, cmd, "cmd should be returned")
	assert.NotNil(t, cmd.RunE, "cmd should be able to be executed")
	assert.Regexp(t, "info", cmd.Use)
	assert.Regexp(t, "pipeline", cmd.Short)
}

func TestInfoCommandRunEClosure(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Client.(*client.MockClient).
		On("FetchPipeline", "foo").
		Return(corev2.FixturePipeline("foo", "default"), nil)
	cli.Config.(*client.MockConfig).On("Format").Return("json")

	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{"foo"})

	assert.NotEmpty(t, out)
	assert.Contains(t, out, "foo")
	assert.Nil(t, err)
}

func TestInfoCommandRunMissingArgs(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Config.(*client.MockConfig).On("Format").Return("json")
	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{})
	require.Error(t, err)
	assert.NotEmpty(t, out)
	assert.Contains(t, out, "Usage")
}

func TestInfoCommandRunEClosureWithTable(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Client.(*client.MockClient).
		On("FetchPipeline", "foo").
		Return(corev2.FixturePipeline("foo", "default"), nil)
	cli.Config.(*client.MockConfig).On("Format").Return("tabular")

	cmd := InfoCommand(cli)
	require.NoError(t, cmd.Flags().Set("format", "tabular"))

	out, err := test.RunCmd(cmd, []string{"foo"})
	require.NoError(t, err)
	assert.NotEmpty(t, out)
	assert.Contains(t, out, "Name")
	assert.Contains(t, out, "Workflows")
}

func TestInfoCommandRunEClosureWithErr(t *testing.T) {
	cli := test.NewMockCLI()
	cli.Client.(*client.MockClient).
		On("FetchPipeline", "foo").
		Return(corev2.FixturePipeline("foo", "default"), fmt.Errorf("error"))
	cli.Config.(*client.MockConfig).On("Format").Return("json")

	cmd := InfoCommand(cli)
	out, err := test.RunCmd(cmd, []string{"foo"})

	assert.NotNil(t, err)
	assert.Equal(t, "error", err.Error())
	assert.Empty(t, out)
}
