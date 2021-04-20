### Developer Notes:
Due to a bug in OpenAPITool:  [Build error if 2 properties have binary format](https://github.com/OpenAPITools/openapi-generator/issues/8165).
The generated code should be patched to prevent a breaking change to the API.

```
diff --git a/pkg/external/api_default.go b/pkg/external/api_default.go
index 646bccb..312f4e0 100644
--- a/pkg/external/api_default.go
+++ b/pkg/external/api_default.go
@@ -322,6 +322,7 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFormFileName string
                localVarFileName     string
                localVarFileBytes    []byte
+               localVarFile         *os.File
                localVarReturnValue  RuntimeComplianceCheck
        )
 
@@ -372,7 +373,6 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFormParams.Add("end_time", parameterToString(localVarOptionals.EndTime.Value(), ""))
        }
        localVarFormFileName = "result_file"
-       var localVarFile *os.File
        if localVarOptionals != nil && localVarOptionals.ResultFile.IsSet() {
                localVarFileOk := false
                localVarFile, localVarFileOk = localVarOptionals.ResultFile.Value().(*os.File)
@@ -385,9 +385,9 @@ func (a *DefaultApiService) AddRuntimeComplianceCheck(ctx _context.Context, chec
                localVarFileBytes = fbs
                localVarFileName = localVarFile.Name()
                localVarFile.Close()
+               localVarFile = nil
        }
        localVarFormFileName = "report_file"
-       var localVarFile *os.File
        if localVarOptionals != nil && localVarOptionals.ReportFile.IsSet() {
                localVarFileOk := false
                localVarFile, localVarFileOk = localVarOptionals.ReportFile.Value().(*os.File)

```