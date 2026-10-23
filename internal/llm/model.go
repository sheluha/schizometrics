package llm

type Question struct {
	Type         string            `json:"type"`
	Instructions string            `json:"instructions"`
	Criteria     map[string]string `json:"criteria,omitempty"`
}

type Request struct {
	Model     string               `json:"model"`
	State     string               `json:"state"`
	Questions map[Feature]Question `json:"questions"`
}

type Answer struct {
	Type          string             `json:"type"`
	Noul          *float64           `json:"noul,omitempty"`
	Choice        string             `json:"choice,omitempty"`
	Confidence    *float64           `json:"confidence,omitempty"`
	Probabilities map[string]float64 `json:"probabilities,omitempty"`
}

type Usage struct {
	Questions      int  `json:"questions"`
	Options        int  `json:"options"`
	StateTokens    int  `json:"state_tokens"`
	StateTruncated bool `json:"state_truncated"`
	UnitsCharged   int  `json:"units_charged"`
}

type Response struct {
	Model   string             `json:"model"`
	Answers map[Feature]Answer `json:"answers"`
	Usage   Usage              `json:"usage"`
}

type Feature string

const (
	FeatureInformationRequest       Feature = "information_request"
	FeatureArgumentSupport          Feature = "argument_support"
	FeatureExpressedUncertainty     Feature = "expressed_uncertainty"
	FeatureConcession               Feature = "concession"
	FeatureHumorIrony               Feature = "humor_irony"
	FeatureIrritation               Feature = "irritation"
	FeaturePersonalAttack           Feature = "personal_attack"
	FeatureProvocativeFraming       Feature = "provocative_framing"
	FeatureCategoricalTone          Feature = "categorical_tone"
	FeatureCondescension            Feature = "condescension"
	FeatureLogicalGap               Feature = "logical_gap"
	FeatureOvergeneralization       Feature = "overgeneralization"
	FeatureContradiction            Feature = "contradiction"
	FeatureCausalOversimplification Feature = "causal_oversimplification"
	FeatureSemanticIncongruity      Feature = "semantic_incongruity"
	FeatureGroupDerogation          Feature = "group_derogation"
	FeatureHarmApproval             Feature = "harm_approval"
	FeatureOpponentDistortion       Feature = "opponent_distortion"
	FeatureObjectionEvasion         Feature = "objection_evasion"
	FeaturePositionRevision         Feature = "position_revision"
	FeatureArgumentRepetition       Feature = "argument_repetition"
	FeatureDetailFixation           Feature = "detail_fixation"
	FeatureLiteralInterpretation    Feature = "literal_interpretation"
	FeatureTopicPersistence         Feature = "topic_persistence"
	FeatureContextInsufficient      Feature = "context_insufficient"
)
