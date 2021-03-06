// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ctxcors

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/corestoreio/csfw/config"
	"github.com/corestoreio/csfw/util/log"
)

// Option defines a function argument for the Cors type to apply options.
type Option func(*Cors)

// WithExposedHeaders indicates which headers are safe to expose to the
// API of a CORS API specification.
func WithExposedHeaders(headers ...string) Option {
	return func(c *Cors) {
		c.exposedHeaders = convert(headers, http.CanonicalHeaderKey)
	}
}

// WithAllowedOrigins is a list of origins a cross-domain request can be executed from.
// If the special "*" value is present in the list, all origins will be allowed.
// An origin may contain a wildcard (*) to replace 0 or more characters
// (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
// Only one wildcard can be used per origin.
// Default value is ["*"]
func WithAllowedOrigins(domains ...string) Option {
	// Note: for origins and methods matching, the spec requires a case-sensitive matching.
	// As it may error prone, we chose to ignore the spec here.
	return func(c *Cors) {

		if len(domains) == 0 {
			// Default is all origins
			c.allowedOriginsAll = true
			return
		}

		c.allowedOrigins = []string{}
		c.allowedWOrigins = []wildcard{}
		for _, origin := range domains {
			// Normalize
			origin = strings.ToLower(origin)
			if origin == "*" {
				// If "*" is present in the list, turn the whole list into a match all
				c.allowedOriginsAll = true
				c.allowedOrigins = nil
				c.allowedWOrigins = nil
				break
			} else if i := strings.IndexByte(origin, '*'); i >= 0 {
				// Split the origin in two: start and end string without the *
				w := wildcard{origin[0:i], origin[i+1:]}
				c.allowedWOrigins = append(c.allowedWOrigins, w)
			} else {
				c.allowedOrigins = append(c.allowedOrigins, origin)
			}
		}
	}
}

// WithAllowOriginFunc convenient helper function.
// AllowOriginFunc is a custom function to validate the origin. It take the origin
// as argument and returns true if allowed or false otherwise. If this option is
// set, the content of AllowedOrigins is ignored.
func WithAllowOriginFunc(f func(origin string) bool) Option {
	return func(c *Cors) {
		c.AllowOriginFunc = f
	}
}

// WithAllowedMethods is a list of methods the client is allowed to use with
// cross-domain requests. Default value is simple methods (GET and POST)
func WithAllowedMethods(methods ...string) Option {
	return func(c *Cors) {
		// Allowed Methods
		// Note: for origins and methods matching, the spec requires a case-sensitive matching.
		// As it may error prone, we chose to ignore the spec here.
		c.allowedMethods = convert(methods, strings.ToUpper)
	}
}

// WithAllowedHeaders is list of non simple headers the client is allowed to use with
// cross-domain requests.
// If the special "*" value is present in the list, all headers will be allowed.
// Default value is [] but "Origin" is always appended to the list.
func WithAllowedHeaders(headers ...string) Option {
	return func(c *Cors) {
		// Origin is always appended as some browsers will always request for this header at preflight
		c.allowedHeaders = convert(append(headers, "Origin"), http.CanonicalHeaderKey)
		for _, h := range headers {
			if h == "*" {
				c.allowedHeadersAll = true
				c.allowedHeaders = nil
				break
			}
		}
	}
}

// WithAllowCredentials convenient helper function.
// AllowCredentials indicates whether the request can include user credentials like
// cookies, HTTP authentication or client side SSL certificates.
func WithAllowCredentials() Option {
	return func(c *Cors) {
		c.AllowCredentials = true
	}
}

// WithMaxAge indicates how long (in seconds) the results of a preflight request
// can be cached
func WithMaxAge(seconds time.Duration) Option {
	s := seconds.Seconds()
	return func(c *Cors) {
		if s > 0 {
			c.maxAge = fmt.Sprintf("%.0f", s)
		}
	}
}

// WithOptionsPassthrough convenient helper function.
// OptionsPassthrough instructs preflight to let other potential next handlers to
// process the OPTIONS method. Turn this on if your application handles OPTIONS.
func WithOptionsPassthrough() Option {
	return func(c *Cors) {
		c.OptionsPassthrough = true
	}
}

// WithLogger convenient helper function.
// Mainly used for debugging.
func WithLogger(l log.Logger) Option {
	return func(c *Cors) {
		c.Log = l
	}
}

// WithConfigGetter adding a Getter to the Cors type allows you to run specific
// Cors configuration for each defined scope ID (default, website or store).
// If you add WithLogger all new created Cors types for a scope will inherit
// from that logger.
func WithConfigGetter(cg config.Getter) Option {
	return func(c *Cors) {
		c.config = cg
	}
}
