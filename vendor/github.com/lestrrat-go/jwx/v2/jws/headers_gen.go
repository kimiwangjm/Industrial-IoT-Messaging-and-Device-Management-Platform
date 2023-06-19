// This file is auto-generated by jws/internal/cmd/genheaders/main.go. DO NOT EDIT

package jws

import (
	"bytes"
	"context"
	"fmt"
	"sort"
	"sync"

	"github.com/lestrrat-go/jwx/v2/cert"
	"github.com/lestrrat-go/jwx/v2/internal/base64"
	"github.com/lestrrat-go/jwx/v2/internal/json"
	"github.com/lestrrat-go/jwx/v2/internal/pool"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

const (
	AlgorithmKey              = "alg"
	ContentTypeKey            = "cty"
	CriticalKey               = "crit"
	JWKKey                    = "jwk"
	JWKSetURLKey              = "jku"
	KeyIDKey                  = "kid"
	TypeKey                   = "typ"
	X509CertChainKey          = "x5c"
	X509CertThumbprintKey     = "x5t"
	X509CertThumbprintS256Key = "x5t#S256"
	X509URLKey                = "x5u"
)

// Headers describe a standard Header set.
type Headers interface {
	json.Marshaler
	json.Unmarshaler
	Algorithm() jwa.SignatureAlgorithm
	ContentType() string
	Critical() []string
	JWK() jwk.Key
	JWKSetURL() string
	KeyID() string
	Type() string
	X509CertChain() *cert.Chain
	X509CertThumbprint() string
	X509CertThumbprintS256() string
	X509URL() string
	Iterate(ctx context.Context) Iterator
	Walk(context.Context, Visitor) error
	AsMap(context.Context) (map[string]interface{}, error)
	Copy(context.Context, Headers) error
	Merge(context.Context, Headers) (Headers, error)
	Get(string) (interface{}, bool)
	Set(string, interface{}) error
	Remove(string) error

	// PrivateParams returns the non-standard elements in the source structure
	// WARNING: DO NOT USE PrivateParams() IF YOU HAVE CONCURRENT CODE ACCESSING THEM.
	// Use AsMap() to get a copy of the entire header instead
	PrivateParams() map[string]interface{}
}

type stdHeaders struct {
	algorithm              *jwa.SignatureAlgorithm // https://tools.ietf.org/html/rfc7515#section-4.1.1
	contentType            *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.10
	critical               []string                // https://tools.ietf.org/html/rfc7515#section-4.1.11
	jwk                    jwk.Key                 // https://tools.ietf.org/html/rfc7515#section-4.1.3
	jwkSetURL              *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.2
	keyID                  *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.4
	typ                    *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.9
	x509CertChain          *cert.Chain             // https://tools.ietf.org/html/rfc7515#section-4.1.6
	x509CertThumbprint     *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.7
	x509CertThumbprintS256 *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.8
	x509URL                *string                 // https://tools.ietf.org/html/rfc7515#section-4.1.5
	privateParams          map[string]interface{}
	mu                     *sync.RWMutex
	dc                     DecodeCtx
	raw                    []byte // stores the raw version of the header so it can be used later
}

func NewHeaders() Headers {
	return &stdHeaders{
		mu: &sync.RWMutex{},
	}
}

func (h *stdHeaders) Algorithm() jwa.SignatureAlgorithm {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.algorithm == nil {
		return ""
	}
	return *(h.algorithm)
}

func (h *stdHeaders) ContentType() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.contentType == nil {
		return ""
	}
	return *(h.contentType)
}

func (h *stdHeaders) Critical() []string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.critical
}

func (h *stdHeaders) JWK() jwk.Key {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.jwk
}

func (h *stdHeaders) JWKSetURL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.jwkSetURL == nil {
		return ""
	}
	return *(h.jwkSetURL)
}

func (h *stdHeaders) KeyID() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.keyID == nil {
		return ""
	}
	return *(h.keyID)
}

func (h *stdHeaders) Type() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.typ == nil {
		return ""
	}
	return *(h.typ)
}

func (h *stdHeaders) X509CertChain() *cert.Chain {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.x509CertChain
}

func (h *stdHeaders) X509CertThumbprint() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprint == nil {
		return ""
	}
	return *(h.x509CertThumbprint)
}

func (h *stdHeaders) X509CertThumbprintS256() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509CertThumbprintS256 == nil {
		return ""
	}
	return *(h.x509CertThumbprintS256)
}

func (h *stdHeaders) X509URL() string {
	h.mu.RLock()
	defer h.mu.RUnlock()
	if h.x509URL == nil {
		return ""
	}
	return *(h.x509URL)
}

func (h *stdHeaders) clear() {
	h.algorithm = nil
	h.contentType = nil
	h.critical = nil
	h.jwk = nil
	h.jwkSetURL = nil
	h.keyID = nil
	h.typ = nil
	h.x509CertChain = nil
	h.x509CertThumbprint = nil
	h.x509CertThumbprintS256 = nil
	h.x509URL = nil
	h.privateParams = nil
	h.raw = nil
}

func (h *stdHeaders) DecodeCtx() DecodeCtx {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.dc
}

func (h *stdHeaders) SetDecodeCtx(dc DecodeCtx) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.dc = dc
}

func (h *stdHeaders) rawBuffer() []byte {
	return h.raw
}

func (h *stdHeaders) makePairs() []*HeaderPair {
	h.mu.RLock()
	defer h.mu.RUnlock()
	var pairs []*HeaderPair
	if h.algorithm != nil {
		pairs = append(pairs, &HeaderPair{Key: AlgorithmKey, Value: *(h.algorithm)})
	}
	if h.contentType != nil {
		pairs = append(pairs, &HeaderPair{Key: ContentTypeKey, Value: *(h.contentType)})
	}
	if h.critical != nil {
		pairs = append(pairs, &HeaderPair{Key: CriticalKey, Value: h.critical})
	}
	if h.jwk != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKKey, Value: h.jwk})
	}
	if h.jwkSetURL != nil {
		pairs = append(pairs, &HeaderPair{Key: JWKSetURLKey, Value: *(h.jwkSetURL)})
	}
	if h.keyID != nil {
		pairs = append(pairs, &HeaderPair{Key: KeyIDKey, Value: *(h.keyID)})
	}
	if h.typ != nil {
		pairs = append(pairs, &HeaderPair{Key: TypeKey, Value: *(h.typ)})
	}
	if h.x509CertChain != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertChainKey, Value: h.x509CertChain})
	}
	if h.x509CertThumbprint != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintKey, Value: *(h.x509CertThumbprint)})
	}
	if h.x509CertThumbprintS256 != nil {
		pairs = append(pairs, &HeaderPair{Key: X509CertThumbprintS256Key, Value: *(h.x509CertThumbprintS256)})
	}
	if h.x509URL != nil {
		pairs = append(pairs, &HeaderPair{Key: X509URLKey, Value: *(h.x509URL)})
	}
	for k, v := range h.privateParams {
		pairs = append(pairs, &HeaderPair{Key: k, Value: v})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Key.(string) < pairs[j].Key.(string)
	})
	return pairs
}

func (h *stdHeaders) PrivateParams() map[string]interface{} {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.privateParams
}

func (h *stdHeaders) Get(name string) (interface{}, bool) {
	h.mu.RLock()
	defer h.mu.RUnlock()
	switch name {
	case AlgorithmKey:
		if h.algorithm == nil {
			return nil, false
		}
		return *(h.algorithm), true
	case ContentTypeKey:
		if h.contentType == nil {
			return nil, false
		}
		return *(h.contentType), true
	case CriticalKey:
		if h.critical == nil {
			return nil, false
		}
		return h.critical, true
	case JWKKey:
		if h.jwk == nil {
			return nil, false
		}
		return h.jwk, true
	case JWKSetURLKey:
		if h.jwkSetURL == nil {
			return nil, false
		}
		return *(h.jwkSetURL), true
	case KeyIDKey:
		if h.keyID == nil {
			return nil, false
		}
		return *(h.keyID), true
	case TypeKey:
		if h.typ == nil {
			return nil, false
		}
		return *(h.typ), true
	case X509CertChainKey:
		if h.x509CertChain == nil {
			return nil, false
		}
		return h.x509CertChain, true
	case X509CertThumbprintKey:
		if h.x509CertThumbprint == nil {
			return nil, false
		}
		return *(h.x509CertThumbprint), true
	case X509CertThumbprintS256Key:
		if h.x509CertThumbprintS256 == nil {
			return nil, false
		}
		return *(h.x509CertThumbprintS256), true
	case X509URLKey:
		if h.x509URL == nil {
			return nil, false
		}
		return *(h.x509URL), true
	default:
		v, ok := h.privateParams[name]
		return v, ok
	}
}

func (h *stdHeaders) Set(name string, value interface{}) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	return h.setNoLock(name, value)
}

func (h *stdHeaders) setNoLock(name string, value interface{}) error {
	switch name {
	case AlgorithmKey:
		var acceptor jwa.SignatureAlgorithm
		if err := acceptor.Accept(value); err != nil {
			return fmt.Errorf(`invalid value for %s key: %w`, AlgorithmKey, err)
		}
		h.algorithm = &acceptor
		return nil
	case ContentTypeKey:
		if v, ok := value.(string); ok {
			h.contentType = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, ContentTypeKey, value)
	case CriticalKey:
		if v, ok := value.([]string); ok {
			h.critical = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, CriticalKey, value)
	case JWKKey:
		if v, ok := value.(jwk.Key); ok {
			h.jwk = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, JWKKey, value)
	case JWKSetURLKey:
		if v, ok := value.(string); ok {
			h.jwkSetURL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, JWKSetURLKey, value)
	case KeyIDKey:
		if v, ok := value.(string); ok {
			h.keyID = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, KeyIDKey, value)
	case TypeKey:
		if v, ok := value.(string); ok {
			h.typ = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, TypeKey, value)
	case X509CertChainKey:
		if v, ok := value.(*cert.Chain); ok {
			h.x509CertChain = v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertChainKey, value)
	case X509CertThumbprintKey:
		if v, ok := value.(string); ok {
			h.x509CertThumbprint = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintKey, value)
	case X509CertThumbprintS256Key:
		if v, ok := value.(string); ok {
			h.x509CertThumbprintS256 = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509CertThumbprintS256Key, value)
	case X509URLKey:
		if v, ok := value.(string); ok {
			h.x509URL = &v
			return nil
		}
		return fmt.Errorf(`invalid value for %s key: %T`, X509URLKey, value)
	default:
		if h.privateParams == nil {
			h.privateParams = map[string]interface{}{}
		}
		h.privateParams[name] = value
	}
	return nil
}

func (h *stdHeaders) Remove(key string) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	switch key {
	case AlgorithmKey:
		h.algorithm = nil
	case ContentTypeKey:
		h.contentType = nil
	case CriticalKey:
		h.critical = nil
	case JWKKey:
		h.jwk = nil
	case JWKSetURLKey:
		h.jwkSetURL = nil
	case KeyIDKey:
		h.keyID = nil
	case TypeKey:
		h.typ = nil
	case X509CertChainKey:
		h.x509CertChain = nil
	case X509CertThumbprintKey:
		h.x509CertThumbprint = nil
	case X509CertThumbprintS256Key:
		h.x509CertThumbprintS256 = nil
	case X509URLKey:
		h.x509URL = nil
	default:
		delete(h.privateParams, key)
	}
	return nil
}

func (h *stdHeaders) UnmarshalJSON(buf []byte) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.clear()
	dec := json.NewDecoder(bytes.NewReader(buf))
LOOP:
	for {
		tok, err := dec.Token()
		if err != nil {
			return fmt.Errorf(`error reading token: %w`, err)
		}
		switch tok := tok.(type) {
		case json.Delim:
			// Assuming we're doing everything correctly, we should ONLY
			// get either '{' or '}' here.
			if tok == '}' { // End of object
				break LOOP
			} else if tok != '{' {
				return fmt.Errorf(`expected '{', but got '%c'`, tok)
			}
		case string: // Objects can only have string keys
			switch tok {
			case AlgorithmKey:
				var decoded jwa.SignatureAlgorithm
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, AlgorithmKey, err)
				}
				h.algorithm = &decoded
			case ContentTypeKey:
				if err := json.AssignNextStringToken(&h.contentType, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, ContentTypeKey, err)
				}
			case CriticalKey:
				var decoded []string
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, CriticalKey, err)
				}
				h.critical = decoded
			case JWKKey:
				var buf json.RawMessage
				if err := dec.Decode(&buf); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, JWKKey, err)
				}
				key, err := jwk.ParseKey(buf)
				if err != nil {
					return fmt.Errorf(`failed to parse JWK for key %s: %w`, JWKKey, err)
				}
				h.jwk = key
			case JWKSetURLKey:
				if err := json.AssignNextStringToken(&h.jwkSetURL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, JWKSetURLKey, err)
				}
			case KeyIDKey:
				if err := json.AssignNextStringToken(&h.keyID, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, KeyIDKey, err)
				}
			case TypeKey:
				if err := json.AssignNextStringToken(&h.typ, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, TypeKey, err)
				}
			case X509CertChainKey:
				var decoded cert.Chain
				if err := dec.Decode(&decoded); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertChainKey, err)
				}
				h.x509CertChain = &decoded
			case X509CertThumbprintKey:
				if err := json.AssignNextStringToken(&h.x509CertThumbprint, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintKey, err)
				}
			case X509CertThumbprintS256Key:
				if err := json.AssignNextStringToken(&h.x509CertThumbprintS256, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509CertThumbprintS256Key, err)
				}
			case X509URLKey:
				if err := json.AssignNextStringToken(&h.x509URL, dec); err != nil {
					return fmt.Errorf(`failed to decode value for key %s: %w`, X509URLKey, err)
				}
			default:
				decoded, err := registry.Decode(dec, tok)
				if err != nil {
					return err
				}
				h.setNoLock(tok, decoded)
			}
		default:
			return fmt.Errorf(`invalid token %T`, tok)
		}
	}
	h.raw = buf
	return nil
}

func (h stdHeaders) MarshalJSON() ([]byte, error) {
	buf := pool.GetBytesBuffer()
	defer pool.ReleaseBytesBuffer(buf)
	buf.WriteByte('{')
	enc := json.NewEncoder(buf)
	for i, p := range h.makePairs() {
		if i > 0 {
			buf.WriteRune(',')
		}
		buf.WriteRune('"')
		buf.WriteString(p.Key.(string))
		buf.WriteString(`":`)
		v := p.Value
		switch v := v.(type) {
		case []byte:
			buf.WriteRune('"')
			buf.WriteString(base64.EncodeToString(v))
			buf.WriteRune('"')
		default:
			if err := enc.Encode(v); err != nil {
				return nil, fmt.Errorf(`failed to encode value for field %s: %w`, p.Key, err)
			}
			buf.Truncate(buf.Len() - 1)
		}
	}
	buf.WriteByte('}')
	ret := make([]byte, buf.Len())
	copy(ret, buf.Bytes())
	return ret, nil
}