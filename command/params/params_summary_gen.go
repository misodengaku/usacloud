// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ShowSummaryParam is input parameters for the sacloud API
type ShowSummaryParam struct {
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	PaidResourcesOnly bool     `json:"paid-resources-only"`
}

// NewShowSummaryParam return new ShowSummaryParam
func NewShowSummaryParam() *ShowSummaryParam {
	return &ShowSummaryParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ShowSummaryParam) FillValueToSkeleton() {
	if isEmpty(p.ParamTemplate) {
		p.ParamTemplate = ""
	}
	if isEmpty(p.ParamTemplateFile) {
		p.ParamTemplateFile = ""
	}
	if isEmpty(p.GenerateSkeleton) {
		p.GenerateSkeleton = false
	}
	if isEmpty(p.OutputType) {
		p.OutputType = ""
	}
	if isEmpty(p.Column) {
		p.Column = []string{""}
	}
	if isEmpty(p.Quiet) {
		p.Quiet = false
	}
	if isEmpty(p.Format) {
		p.Format = ""
	}
	if isEmpty(p.FormatFile) {
		p.FormatFile = ""
	}
	if isEmpty(p.PaidResourcesOnly) {
		p.PaidResourcesOnly = false
	}

}

// Validate checks current values in model
func (p *ShowSummaryParam) Validate() []error {
	errors := []error{}

	{
		validator := schema.ValidateInStrValues("json", "csv", "tsv")
		errs := validator("--output-type", p.OutputType)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateInputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateOutputOption(p)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

	return errors
}

func (p *ShowSummaryParam) GetResourceDef() *schema.Resource {
	return define.Resources["Summary"]
}

func (p *ShowSummaryParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["show"]
}

func (p *ShowSummaryParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ShowSummaryParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ShowSummaryParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ShowSummaryParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ShowSummaryParam) GetOutputFormat() string {
	return "table"
}

func (p *ShowSummaryParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ShowSummaryParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ShowSummaryParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ShowSummaryParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ShowSummaryParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ShowSummaryParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ShowSummaryParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ShowSummaryParam) GetOutputType() string {
	return p.OutputType
}
func (p *ShowSummaryParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ShowSummaryParam) GetColumn() []string {
	return p.Column
}
func (p *ShowSummaryParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ShowSummaryParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ShowSummaryParam) SetFormat(v string) {
	p.Format = v
}

func (p *ShowSummaryParam) GetFormat() string {
	return p.Format
}
func (p *ShowSummaryParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ShowSummaryParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ShowSummaryParam) SetPaidResourcesOnly(v bool) {
	p.PaidResourcesOnly = v
}

func (p *ShowSummaryParam) GetPaidResourcesOnly() bool {
	return p.PaidResourcesOnly
}
