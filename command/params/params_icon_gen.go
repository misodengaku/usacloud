// Code generated by 'github.com/sacloud/usacloud/tools/gen-input-models'; DO NOT EDIT

package params

import (
	"github.com/sacloud/usacloud/define"
	"github.com/sacloud/usacloud/output"
	"github.com/sacloud/usacloud/schema"
)

// ListIconParam is input parameters for the sacloud API
type ListIconParam struct {
	Name              []string `json:"name"`
	Id                []int64  `json:"id"`
	Scope             string   `json:"scope"`
	Tags              []string `json:"tags"`
	From              int      `json:"from"`
	Max               int      `json:"max"`
	Sort              []string `json:"sort"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
}

// NewListIconParam return new ListIconParam
func NewListIconParam() *ListIconParam {
	return &ListIconParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ListIconParam) FillValueToSkeleton() {
	if isEmpty(p.Name) {
		p.Name = []string{""}
	}
	if isEmpty(p.Id) {
		p.Id = []int64{0}
	}
	if isEmpty(p.Scope) {
		p.Scope = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.From) {
		p.From = 0
	}
	if isEmpty(p.Max) {
		p.Max = 0
	}
	if isEmpty(p.Sort) {
		p.Sort = []string{""}
	}
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

}

// Validate checks current values in model
func (p *ListIconParam) Validate() []error {
	errors := []error{}
	{
		errs := validateConflicts("--name", p.Name, map[string]interface{}{

			"--id": p.Id,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["list"].Params["id"].ValidateFunc
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		errs := validateConflicts("--id", p.Id, map[string]interface{}{

			"--name": p.Name,
		})
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["list"].Params["scope"].ValidateFunc
		errs := validator("--scope", p.Scope)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["list"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

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

func (p *ListIconParam) GetResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *ListIconParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["list"]
}

func (p *ListIconParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ListIconParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ListIconParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ListIconParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ListIconParam) GetOutputFormat() string {
	return "table"
}

func (p *ListIconParam) SetName(v []string) {
	p.Name = v
}

func (p *ListIconParam) GetName() []string {
	return p.Name
}
func (p *ListIconParam) SetId(v []int64) {
	p.Id = v
}

func (p *ListIconParam) GetId() []int64 {
	return p.Id
}
func (p *ListIconParam) SetScope(v string) {
	p.Scope = v
}

func (p *ListIconParam) GetScope() string {
	return p.Scope
}
func (p *ListIconParam) SetTags(v []string) {
	p.Tags = v
}

func (p *ListIconParam) GetTags() []string {
	return p.Tags
}
func (p *ListIconParam) SetFrom(v int) {
	p.From = v
}

func (p *ListIconParam) GetFrom() int {
	return p.From
}
func (p *ListIconParam) SetMax(v int) {
	p.Max = v
}

func (p *ListIconParam) GetMax() int {
	return p.Max
}
func (p *ListIconParam) SetSort(v []string) {
	p.Sort = v
}

func (p *ListIconParam) GetSort() []string {
	return p.Sort
}
func (p *ListIconParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ListIconParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ListIconParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ListIconParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ListIconParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ListIconParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ListIconParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ListIconParam) GetOutputType() string {
	return p.OutputType
}
func (p *ListIconParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ListIconParam) GetColumn() []string {
	return p.Column
}
func (p *ListIconParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ListIconParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ListIconParam) SetFormat(v string) {
	p.Format = v
}

func (p *ListIconParam) GetFormat() string {
	return p.Format
}
func (p *ListIconParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ListIconParam) GetFormatFile() string {
	return p.FormatFile
}

// CreateIconParam is input parameters for the sacloud API
type CreateIconParam struct {
	Image             string   `json:"image"`
	Name              string   `json:"name"`
	Tags              []string `json:"tags"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
}

// NewCreateIconParam return new CreateIconParam
func NewCreateIconParam() *CreateIconParam {
	return &CreateIconParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *CreateIconParam) FillValueToSkeleton() {
	if isEmpty(p.Image) {
		p.Image = ""
	}
	if isEmpty(p.Name) {
		p.Name = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
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

}

// Validate checks current values in model
func (p *CreateIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateRequired
		errs := validator("--image", p.Image)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["image"].ValidateFunc
		errs := validator("--image", p.Image)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateRequired
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["create"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

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

func (p *CreateIconParam) GetResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *CreateIconParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["create"]
}

func (p *CreateIconParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *CreateIconParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *CreateIconParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *CreateIconParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *CreateIconParam) GetOutputFormat() string {
	return "table"
}

func (p *CreateIconParam) SetImage(v string) {
	p.Image = v
}

func (p *CreateIconParam) GetImage() string {
	return p.Image
}
func (p *CreateIconParam) SetName(v string) {
	p.Name = v
}

func (p *CreateIconParam) GetName() string {
	return p.Name
}
func (p *CreateIconParam) SetTags(v []string) {
	p.Tags = v
}

func (p *CreateIconParam) GetTags() []string {
	return p.Tags
}
func (p *CreateIconParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *CreateIconParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *CreateIconParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *CreateIconParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *CreateIconParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *CreateIconParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *CreateIconParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *CreateIconParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *CreateIconParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *CreateIconParam) GetOutputType() string {
	return p.OutputType
}
func (p *CreateIconParam) SetColumn(v []string) {
	p.Column = v
}

func (p *CreateIconParam) GetColumn() []string {
	return p.Column
}
func (p *CreateIconParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *CreateIconParam) GetQuiet() bool {
	return p.Quiet
}
func (p *CreateIconParam) SetFormat(v string) {
	p.Format = v
}

func (p *CreateIconParam) GetFormat() string {
	return p.Format
}
func (p *CreateIconParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *CreateIconParam) GetFormatFile() string {
	return p.FormatFile
}

// ReadIconParam is input parameters for the sacloud API
type ReadIconParam struct {
	Selector          []string `json:"selector"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Id                int64    `json:"id"`
}

// NewReadIconParam return new ReadIconParam
func NewReadIconParam() *ReadIconParam {
	return &ReadIconParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *ReadIconParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
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
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *ReadIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

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

func (p *ReadIconParam) GetResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *ReadIconParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["read"]
}

func (p *ReadIconParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *ReadIconParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *ReadIconParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *ReadIconParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *ReadIconParam) GetOutputFormat() string {
	return "table"
}

func (p *ReadIconParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *ReadIconParam) GetSelector() []string {
	return p.Selector
}
func (p *ReadIconParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *ReadIconParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *ReadIconParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *ReadIconParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *ReadIconParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *ReadIconParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *ReadIconParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *ReadIconParam) GetOutputType() string {
	return p.OutputType
}
func (p *ReadIconParam) SetColumn(v []string) {
	p.Column = v
}

func (p *ReadIconParam) GetColumn() []string {
	return p.Column
}
func (p *ReadIconParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *ReadIconParam) GetQuiet() bool {
	return p.Quiet
}
func (p *ReadIconParam) SetFormat(v string) {
	p.Format = v
}

func (p *ReadIconParam) GetFormat() string {
	return p.Format
}
func (p *ReadIconParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *ReadIconParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *ReadIconParam) SetId(v int64) {
	p.Id = v
}

func (p *ReadIconParam) GetId() int64 {
	return p.Id
}

// UpdateIconParam is input parameters for the sacloud API
type UpdateIconParam struct {
	Selector          []string `json:"selector"`
	Name              string   `json:"name"`
	Tags              []string `json:"tags"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Id                int64    `json:"id"`
}

// NewUpdateIconParam return new UpdateIconParam
func NewUpdateIconParam() *UpdateIconParam {
	return &UpdateIconParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *UpdateIconParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Name) {
		p.Name = ""
	}
	if isEmpty(p.Tags) {
		p.Tags = []string{""}
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
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
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *UpdateIconParam) Validate() []error {
	errors := []error{}
	{
		validator := define.Resources["Icon"].Commands["update"].Params["name"].ValidateFunc
		errs := validator("--name", p.Name)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := define.Resources["Icon"].Commands["update"].Params["tags"].ValidateFunc
		errs := validator("--tags", p.Tags)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

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

func (p *UpdateIconParam) GetResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *UpdateIconParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["update"]
}

func (p *UpdateIconParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *UpdateIconParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *UpdateIconParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *UpdateIconParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *UpdateIconParam) GetOutputFormat() string {
	return "table"
}

func (p *UpdateIconParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *UpdateIconParam) GetSelector() []string {
	return p.Selector
}
func (p *UpdateIconParam) SetName(v string) {
	p.Name = v
}

func (p *UpdateIconParam) GetName() string {
	return p.Name
}
func (p *UpdateIconParam) SetTags(v []string) {
	p.Tags = v
}

func (p *UpdateIconParam) GetTags() []string {
	return p.Tags
}
func (p *UpdateIconParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *UpdateIconParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *UpdateIconParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *UpdateIconParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *UpdateIconParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *UpdateIconParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *UpdateIconParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *UpdateIconParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *UpdateIconParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *UpdateIconParam) GetOutputType() string {
	return p.OutputType
}
func (p *UpdateIconParam) SetColumn(v []string) {
	p.Column = v
}

func (p *UpdateIconParam) GetColumn() []string {
	return p.Column
}
func (p *UpdateIconParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *UpdateIconParam) GetQuiet() bool {
	return p.Quiet
}
func (p *UpdateIconParam) SetFormat(v string) {
	p.Format = v
}

func (p *UpdateIconParam) GetFormat() string {
	return p.Format
}
func (p *UpdateIconParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *UpdateIconParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *UpdateIconParam) SetId(v int64) {
	p.Id = v
}

func (p *UpdateIconParam) GetId() int64 {
	return p.Id
}

// DeleteIconParam is input parameters for the sacloud API
type DeleteIconParam struct {
	Selector          []string `json:"selector"`
	Assumeyes         bool     `json:"assumeyes"`
	ParamTemplate     string   `json:"param-template"`
	ParamTemplateFile string   `json:"param-template-file"`
	GenerateSkeleton  bool     `json:"generate-skeleton"`
	OutputType        string   `json:"output-type"`
	Column            []string `json:"column"`
	Quiet             bool     `json:"quiet"`
	Format            string   `json:"format"`
	FormatFile        string   `json:"format-file"`
	Id                int64    `json:"id"`
}

// NewDeleteIconParam return new DeleteIconParam
func NewDeleteIconParam() *DeleteIconParam {
	return &DeleteIconParam{}
}

// FillValueToSkeleton fill values to empty fields
func (p *DeleteIconParam) FillValueToSkeleton() {
	if isEmpty(p.Selector) {
		p.Selector = []string{""}
	}
	if isEmpty(p.Assumeyes) {
		p.Assumeyes = false
	}
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
	if isEmpty(p.Id) {
		p.Id = 0
	}

}

// Validate checks current values in model
func (p *DeleteIconParam) Validate() []error {
	errors := []error{}
	{
		validator := validateSakuraID
		errs := validator("--id", p.Id)
		if errs != nil {
			errors = append(errors, errs...)
		}
	}

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

func (p *DeleteIconParam) GetResourceDef() *schema.Resource {
	return define.Resources["Icon"]
}

func (p *DeleteIconParam) GetCommandDef() *schema.Command {
	return p.GetResourceDef().Commands["delete"]
}

func (p *DeleteIconParam) GetIncludeFields() []string {
	return p.GetCommandDef().IncludeFields
}

func (p *DeleteIconParam) GetExcludeFields() []string {
	return p.GetCommandDef().ExcludeFields
}

func (p *DeleteIconParam) GetTableType() output.TableType {
	return p.GetCommandDef().TableType
}

func (p *DeleteIconParam) GetColumnDefs() []output.ColumnDef {
	return p.GetCommandDef().TableColumnDefines
}

func (p *DeleteIconParam) GetOutputFormat() string {
	return "table"
}

func (p *DeleteIconParam) SetSelector(v []string) {
	p.Selector = v
}

func (p *DeleteIconParam) GetSelector() []string {
	return p.Selector
}
func (p *DeleteIconParam) SetAssumeyes(v bool) {
	p.Assumeyes = v
}

func (p *DeleteIconParam) GetAssumeyes() bool {
	return p.Assumeyes
}
func (p *DeleteIconParam) SetParamTemplate(v string) {
	p.ParamTemplate = v
}

func (p *DeleteIconParam) GetParamTemplate() string {
	return p.ParamTemplate
}
func (p *DeleteIconParam) SetParamTemplateFile(v string) {
	p.ParamTemplateFile = v
}

func (p *DeleteIconParam) GetParamTemplateFile() string {
	return p.ParamTemplateFile
}
func (p *DeleteIconParam) SetGenerateSkeleton(v bool) {
	p.GenerateSkeleton = v
}

func (p *DeleteIconParam) GetGenerateSkeleton() bool {
	return p.GenerateSkeleton
}
func (p *DeleteIconParam) SetOutputType(v string) {
	p.OutputType = v
}

func (p *DeleteIconParam) GetOutputType() string {
	return p.OutputType
}
func (p *DeleteIconParam) SetColumn(v []string) {
	p.Column = v
}

func (p *DeleteIconParam) GetColumn() []string {
	return p.Column
}
func (p *DeleteIconParam) SetQuiet(v bool) {
	p.Quiet = v
}

func (p *DeleteIconParam) GetQuiet() bool {
	return p.Quiet
}
func (p *DeleteIconParam) SetFormat(v string) {
	p.Format = v
}

func (p *DeleteIconParam) GetFormat() string {
	return p.Format
}
func (p *DeleteIconParam) SetFormatFile(v string) {
	p.FormatFile = v
}

func (p *DeleteIconParam) GetFormatFile() string {
	return p.FormatFile
}
func (p *DeleteIconParam) SetId(v int64) {
	p.Id = v
}

func (p *DeleteIconParam) GetId() int64 {
	return p.Id
}
