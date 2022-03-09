package web

type MatchType string

var MatchTypes = struct {
	Begins   MatchType
	Contains MatchType
	Ends     MatchType
}{
	"Begins",
	"Contains",
	"Ends",
}

type ApplicationType string

var ApplicationTypes = struct {
	AutoInjected             ApplicationType
	BrowserExtensionInjected ApplicationType
	ManuallyInjected         ApplicationType
}{
	"AUTO_INJECTED",
	"BROWSER_EXTENSION_INJECTED",
	"MANUALLY_INJECTED",
}

type Comparator string

var Comparators = struct {
	Equals             Comparator
	GreaterThanOrEqual Comparator
	LowerThanOrEqual   Comparator
}{
	"EQUALS",
	"GREATER_THAN_OR_EQUAL",
	"LOWER_THAN_OR_EQUAL",
}

type Platform string

var Platforms = struct {
	All     Platform
	Desktop Platform
	Mobile  Platform
}{
	"ALL",
	"DESKTOP",
	"MOBILE",
}

type BrowserType string

var BrowserTypes = struct {
	AndroidWebKit    BrowserType
	BotsSpiders      BrowserType
	Chrome           BrowserType
	Edge             BrowserType
	Firefox          BrowserType
	InternetExplorer BrowserType
	Opera            BrowserType
	Safari           BrowserType
}{
	"ANDROID_WEBKIT",
	"BOTS_SPIDERS",
	"CHROME",
	"EDGE",
	"FIREFOX",
	"INTERNET_EXPLORER",
	"OPERA",
	"SAFARI",
}

type RestrictionMode string

var RestrictionModes = struct {
	Exclude RestrictionMode
	Include RestrictionMode
}{
	"EXCLUDE",
	"INCLUDE",
}

type ConversionGoalType string

var ConversionGoalTypes = struct {
	Destination     ConversionGoalType
	UserAction      ConversionGoalType
	VisitDuration   ConversionGoalType
	VisitNumActions ConversionGoalType
}{
	"Destination",
	"UserAction",
	"VisitDuration",
	"VisitNumActions",
}

type InjectionMode string

var InjectionModes = struct {
	CodeSnippet      InjectionMode
	CodeSnippetAsync InjectionMode
	InlineCode       InjectionMode
	JavaScriptTag    InjectionMode
}{
	"CODE_SNIPPET",
	"CODE_SNIPPET_ASYNC",
	"INLINE_CODE",
	"JAVASCRIPT_TAG",
}

type InjectionTarget string

var InjectionTargets = struct {
	PageQuery InjectionTarget
	URL       InjectionTarget
}{
	"PAGE_QUERY",
	"URL",
}

type URLOperator string

var URLOperators = struct {
	AllPages   URLOperator
	Contains   URLOperator
	EndsWith   URLOperator
	Equals     URLOperator
	StartsWith URLOperator
}{
	"ALL_PAGES",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
	"STARTS_WITH",
}

type JSInjectionRule string

var JSInjectionRules = struct {
	AfterSpecificHTML  JSInjectionRule
	AutomaticInjection JSInjectionRule
	BeforeSpecificHTML JSInjectionRule
	DoNotInject        JSInjectionRule
}{
	"AFTER_SPECIFIC_HTML",
	"AUTOMATIC_INJECTION",
	"BEFORE_SPECIFIC_HTML",
	"DO_NOT_INJECT",
}

type LoadActionKeyPerformanceMetric string

var LoadActionKeyPerformanceMetrics = struct {
	ActionDuration         LoadActionKeyPerformanceMetric
	CumulativeLayoutShift  LoadActionKeyPerformanceMetric
	DomInteractive         LoadActionKeyPerformanceMetric
	FirstInputDelay        LoadActionKeyPerformanceMetric
	LargestContentfulPaint LoadActionKeyPerformanceMetric
	LoadEventEnd           LoadActionKeyPerformanceMetric
	LoadEventStart         LoadActionKeyPerformanceMetric
	ResponseEnd            LoadActionKeyPerformanceMetric
	ResponseStart          LoadActionKeyPerformanceMetric
	SpeedIndex             LoadActionKeyPerformanceMetric
	VisuallyComplete       LoadActionKeyPerformanceMetric
}{
	"ACTION_DURATION",
	"CUMULATIVE_LAYOUT_SHIFT",
	"DOM_INTERACTIVE",
	"FIRST_INPUT_DELAY",
	"LARGEST_CONTENTFUL_PAINT",
	"LOAD_EVENT_END",
	"LOAD_EVENT_START",
	"RESPONSE_END",
	"RESPONSE_START",
	"SPEED_INDEX",
	"VISUALLY_COMPLETE",
}

type ResourceTimingCaptureType string

var ResourceTimingCaptureTypes = struct {
	CaptureAllSummaries     ResourceTimingCaptureType
	CaptureFullDetails      ResourceTimingCaptureType
	CaptureLimitedSummaries ResourceTimingCaptureType
}{
	"CAPTURE_ALL_SUMMARIES",
	"CAPTURE_FULL_DETAILS",
	"CAPTURE_LIMITED_SUMMARIES",
}

type Operator string

var Operators = struct {
	Contains                    Operator
	EndsWith                    Operator
	Equals                      Operator
	IsEmpty                     Operator
	IsNotEmpty                  Operator
	MatchesRegularExpression    Operator
	NotContains                 Operator
	NotEndsWith                 Operator
	NotEquals                   Operator
	NotMatchesRegularExpression Operator
	NotStartsWith               Operator
	StartsWith                  Operator
}{
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
	"IS_EMPTY",
	"IS_NOT_EMPTY",
	"MATCHES_REGULAR_EXPRESSION",
	"NOT_CONTAINS",
	"NOT_ENDS_WITH",
	"NOT_EQUALS",
	"NOT_MATCHES_REGULAR_EXPRESSION",
	"NOT_STARTS_WITH",
	"STARTS_WITH",
}

type XHRActionKeyPerformanceMetric string

var XHRActionKeyPerformanceMetrics = struct {
	ActionDuration   XHRActionKeyPerformanceMetric
	ResponseEnd      XHRActionKeyPerformanceMetric
	ResponseStart    XHRActionKeyPerformanceMetric
	VisuallyComplete XHRActionKeyPerformanceMetric
}{
	"ACTION_DURATION",
	"RESPONSE_END",
	"RESPONSE_START",
	"VISUALLY_COMPLETE",
}

type MetaDataCapturingType string

var MetaDataCapturingTypes = struct {
	Cookie             MetaDataCapturingType
	CSSSelector        MetaDataCapturingType
	JavaScriptFunction MetaDataCapturingType
	JavaScriptVariable MetaDataCapturingType
	MetaTag            MetaDataCapturingType
	QueryString        MetaDataCapturingType
}{
	"COOKIE",
	"CSS_SELECTOR",
	"JAVA_SCRIPT_FUNCTION",
	"JAVA_SCRIPT_VARIABLE",
	"META_TAG",
	"QUERY_STRING",
}

type PropertyType string

var PropertyTypes = struct {
	Date       PropertyType
	Double     PropertyType
	Long       PropertyType
	LongString PropertyType
	String     PropertyType
}{
	"DATE",
	"DOUBLE",
	"LONG",
	"LONG_STRING",
	"STRING",
}

type PropertyOrigin string

var PropertyOrigins = struct {
	JavaScriptAPI              PropertyOrigin
	MetaData                   PropertyOrigin
	ServerSideRequestAttribute PropertyOrigin
}{
	"JAVASCRIPT_API",
	"META_DATA",
	"SERVER_SIDE_REQUEST_ATTRIBUTE",
}

type Aggregation string

var Aggregations = struct {
	Average Aggregation
	First   Aggregation
	Last    Aggregation
	Maximum Aggregation
	Minimum Aggregation
	Sum     Aggregation
}{
	"AVERAGE",
	"FIRST",
	"LAST",
	"MAXIMUM",
	"MINIMUM",
	"SUM",
}

type MatchEntity string

var MatchEntities = struct {
	ActionName         MatchEntity
	CSSSelector        MatchEntity
	JavaScriptVariable MatchEntity
	MetaTag            MatchEntity
	PagePath           MatchEntity
	PageTitle          MatchEntity
	PageURL            MatchEntity
	URLAnchor          MatchEntity
	XHRURL             MatchEntity
}{
	"ActionName",
	"CssSelector",
	"JavaScriptVariable",
	"MetaTag",
	"PagePath",
	"PageTitle",
	"PageUrl",
	"UrlAnchor",
	"XhrUrl",
}

type ActionType string

var ActionTypes = struct {
	Custom ActionType
	Load   ActionType
	XHR    ActionType
}{
	"Custom",
	"Load",
	"Xhr",
}

type Input string

var Inputs = struct {
	ElementIdentifier Input
	InputType         Input
	MetaData          Input
	PageTitle         Input
	PageURL           Input
	SourceURL         Input
	TopXHRURL         Input
	XHRURL            Input
}{
	"ELEMENT_IDENTIFIER",
	"INPUT_TYPE",
	"METADATA",
	"PAGE_TITLE",
	"PAGE_URL",
	"SOURCE_URL",
	"TOP_XHR_URL",
	"XHR_URL",
}

type ProcessingPart string

var ProcessingParts = struct {
	All    ProcessingPart
	Anchor ProcessingPart
	Path   ProcessingPart
}{
	"ALL",
	"ANCHOR",
	"PATH",
}

type ProcessingStepType string

var ProcessingStepTypes = struct {
	ExtractByRegularExpression   ProcessingStepType
	Replacement                  ProcessingStepType
	ReplaceIDs                   ProcessingStepType
	ReplaceWithPattern           ProcessingStepType
	ReplaceWithRegularExpression ProcessingStepType
	SubString                    ProcessingStepType
}{
	"EXTRACT_BY_REGULAR_EXPRESSION",
	"REPLACEMENT",
	"REPLACE_IDS",
	"REPLACE_WITH_PATTERN",
	"REPLACE_WITH_REGULAR_EXPRESSION",
	"SUBSTRING",
}

type PatternSearchType string

var PatternSearchTypes = struct {
	First PatternSearchType
	Last  PatternSearchType
}{
	"FIRST",
	"LAST",
}

type DoNotTrackBehaviour string

var DoNotTrackBehaviours = struct {
	CaptureAnonymized DoNotTrackBehaviour
	NoNotCapture      DoNotTrackBehaviour
	IgnoreDoNotTrack  DoNotTrackBehaviour
}{
	"CAPTURE_ANONYMIZED",
	"DO_NOT_CAPTURE",
	"IGNORE_DO_NOT_TRACK",
}

type HTTPErrorRuleFilter string

var HTTPErrorRuleFilters = struct {
	BeginsWith HTTPErrorRuleFilter
	Contains   HTTPErrorRuleFilter
	EndsWith   HTTPErrorRuleFilter
	Equals     HTTPErrorRuleFilter
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
}

type CustomErrorRuleKeyMatcher string

var CustomErrorRuleKeyMatchers = struct {
	BeginsWith CustomErrorRuleKeyMatcher
	Contains   CustomErrorRuleKeyMatcher
	EndsWith   CustomErrorRuleKeyMatcher
	Equals     CustomErrorRuleKeyMatcher
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
}

type CustomErrorRuleValueMatcher string

var CustomErrorRuleValueMatchers = struct {
	BeginsWith CustomErrorRuleValueMatcher
	Contains   CustomErrorRuleValueMatcher
	EndsWith   CustomErrorRuleValueMatcher
	Equals     CustomErrorRuleValueMatcher
}{
	"BEGINS_WITH",
	"CONTAINS",
	"ENDS_WITH",
	"EQUALS",
}

type KeyUserActionType string

var KeyUserActionTypes = struct {
	Custom KeyUserActionType
	Load   KeyUserActionType
	XHR    KeyUserActionType
}{
	"Custom",
	"Load",
	"Xhr",
}
