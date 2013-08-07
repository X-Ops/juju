// Copyright 2012, 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package main

import (
	"bytes"
	"fmt"
	. "launchpad.net/gocheck"
	"launchpad.net/juju-core/cmd"
	"launchpad.net/juju-core/juju/testing"
	coretesting "launchpad.net/juju-core/testing"
)

type EndpointSuite struct {
	testing.JujuConnSuite
}

var _ = Suite(&EndpointSuite{})


func (s *EndpointSuite) TestEndpoint(c *C) {
	ctx := coretesting.Context(c)
	code := cmd.Main(&EndpointCommand{}, ctx, []string{})
	c.Check(code, Equals, 0)
	c.Assert(ctx.Stderr.(*bytes.Buffer).String(), Equals, "")
	output := string(ctx.Stdout.(*bytes.Buffer).Bytes())
	info := s.APIInfo(c)
	c.Assert(output, Equals, fmt.Sprintf("%s\n", info.Addrs[0]))
}
