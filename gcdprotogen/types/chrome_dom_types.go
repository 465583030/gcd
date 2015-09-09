// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains the DOM types.
// API Version: 1.1
package types

// Unique DOM node identifier.
type ChromeDOMNodeId int

// Unique DOM node identifier used to reference a node that may not have been pushed to the front-end.
type ChromeDOMBackendNodeId int

// Backend node with a friendly name.
type ChromeDOMBackendNode struct {
	NodeType      int                     `json:"nodeType"` // <code>Node</code>'s nodeType.
	NodeName      string                  `json:"nodeName"` // <code>Node</code>'s nodeName.
	BackendNodeId *ChromeDOMBackendNodeId `json:"backendNodeId"`
}

// Pseudo element type.
type ChromeDOMPseudoType string // possible values: first-line, first-letter, before, after, backdrop, selection, first-line-inherited, scrollbar, scrollbar-thumb, scrollbar-button, scrollbar-track, scrollbar-track-piece, scrollbar-corner, resizer, input-list-button

// Shadow root type.
type ChromeDOMShadowRootType string // possible values: user-agent, open, closed

// DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type ChromeDOMNode struct {
	NodeId           *ChromeDOMNodeId         `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the <code>nodeId</code>. Backend will only push node with given <code>id</code> once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	NodeType         int                      `json:"nodeType"`                   // <code>Node</code>'s nodeType.
	NodeName         string                   `json:"nodeName"`                   // <code>Node</code>'s nodeName.
	LocalName        string                   `json:"localName"`                  // <code>Node</code>'s localName.
	NodeValue        string                   `json:"nodeValue"`                  // <code>Node</code>'s nodeValue.
	ChildNodeCount   int                      `json:"childNodeCount,omitempty"`   // Child count for <code>Container</code> nodes.
	Children         []*ChromeDOMNode         `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string                 `json:"attributes,omitempty"`       // Attributes of the <code>Element</code> node in the form of flat array <code>[name1, value1, name2, value2]</code>.
	DocumentURL      string                   `json:"documentURL,omitempty"`      // Document URL that <code>Document</code> or <code>FrameOwner</code> node points to.
	BaseURL          string                   `json:"baseURL,omitempty"`          // Base URL that <code>Document</code> or <code>FrameOwner</code> node uses for URL completion.
	PublicId         string                   `json:"publicId,omitempty"`         // <code>DocumentType</code>'s publicId.
	SystemId         string                   `json:"systemId,omitempty"`         // <code>DocumentType</code>'s systemId.
	InternalSubset   string                   `json:"internalSubset,omitempty"`   // <code>DocumentType</code>'s internalSubset.
	XmlVersion       string                   `json:"xmlVersion,omitempty"`       // <code>Document</code>'s XML version in case of XML documents.
	Name             string                   `json:"name,omitempty"`             // <code>Attr</code>'s name.
	Value            string                   `json:"value,omitempty"`            // <code>Attr</code>'s value.
	PseudoType       *ChromeDOMPseudoType     `json:"pseudoType,omitempty"`       // Pseudo element type for this node.
	ShadowRootType   *ChromeDOMShadowRootType `json:"shadowRootType,omitempty"`   // Shadow root type.
	FrameId          *ChromePageFrameId       `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *ChromeDOMNode           `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []*ChromeDOMNode         `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *ChromeDOMNode           `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []*ChromeDOMNode         `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *ChromeDOMNode           `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []*ChromeDOMBackendNode  `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
}

// A structure holding an RGBA color.
type ChromeDOMRGBA struct {
	R int     `json:"r"`           // The red component, in the [0-255] range.
	G int     `json:"g"`           // The green component, in the [0-255] range.
	B int     `json:"b"`           // The blue component, in the [0-255] range.
	A float64 `json:"a,omitempty"` // The alpha component, in the [0-1] range (default: 1).
}

// An array of quad vertices, x immediately followed by y for each point, points clock-wise.
type ChromeDOMQuad []float64

// Box model.
type ChromeDOMBoxModel struct {
	Content      *ChromeDOMQuad             `json:"content"`                // Content box
	Padding      *ChromeDOMQuad             `json:"padding"`                // Padding box
	Border       *ChromeDOMQuad             `json:"border"`                 // Border box
	Margin       *ChromeDOMQuad             `json:"margin"`                 // Margin box
	Width        int                        `json:"width"`                  // Node width
	Height       int                        `json:"height"`                 // Node height
	ShapeOutside *ChromeDOMShapeOutsideInfo `json:"shapeOutside,omitempty"` // Shape outside coordinates
}

// CSS Shape Outside details.
type ChromeDOMShapeOutsideInfo struct {
	Bounds      *ChromeDOMQuad `json:"bounds"`      // Shape bounds
	Shape       []string       `json:"shape"`       // Shape coordinate details
	MarginShape []string       `json:"marginShape"` // Margin shape bounds
}

// Rectangle.
type ChromeDOMRect struct {
	X      float64 `json:"x"`      // X coordinate
	Y      float64 `json:"y"`      // Y coordinate
	Width  float64 `json:"width"`  // Rectangle width
	Height float64 `json:"height"` // Rectangle height
}

// Configuration data for the highlighting of page elements.
type ChromeDOMHighlightConfig struct {
	ShowInfo           bool           `json:"showInfo,omitempty"`           // Whether the node info tooltip should be shown (default: false).
	ShowRulers         bool           `json:"showRulers,omitempty"`         // Whether the rulers should be shown (default: false).
	ShowExtensionLines bool           `json:"showExtensionLines,omitempty"` // Whether the extension lines from node to the rulers should be shown (default: false).
	DisplayAsMaterial  bool           `json:"displayAsMaterial,omitempty"`
	ContentColor       *ChromeDOMRGBA `json:"contentColor,omitempty"`     // The content box highlight fill color (default: transparent).
	PaddingColor       *ChromeDOMRGBA `json:"paddingColor,omitempty"`     // The padding highlight fill color (default: transparent).
	BorderColor        *ChromeDOMRGBA `json:"borderColor,omitempty"`      // The border highlight fill color (default: transparent).
	MarginColor        *ChromeDOMRGBA `json:"marginColor,omitempty"`      // The margin highlight fill color (default: transparent).
	EventTargetColor   *ChromeDOMRGBA `json:"eventTargetColor,omitempty"` // The event target element highlight fill color (default: transparent).
	ShapeColor         *ChromeDOMRGBA `json:"shapeColor,omitempty"`       // The shape outside fill color (default: transparent).
	ShapeMarginColor   *ChromeDOMRGBA `json:"shapeMarginColor,omitempty"` // The shape margin fill color (default: transparent).
}

//
type ChromeDOMInspectMode string // possible values: searchForNode, searchForUAShadowDOM, showLayoutEditor, none
