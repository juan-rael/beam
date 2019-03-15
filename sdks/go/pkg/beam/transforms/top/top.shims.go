// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by starcgen. DO NOT EDIT.
// File: top.shims.go

package top

import (
	"reflect"

	// Library imports
	"github.com/apache/beam/sdks/go/pkg/beam/core/runtime"
	"github.com/apache/beam/sdks/go/pkg/beam/core/typex"
	"github.com/apache/beam/sdks/go/pkg/beam/core/util/reflectx"
)

func init() {
	runtime.RegisterType(reflect.TypeOf((*accum)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*combineFn)(nil)).Elem())
	reflectx.RegisterStructWrapper(reflect.TypeOf((*combineFn)(nil)).Elem(), wrapMakerCombineFn)
	reflectx.RegisterFunc(reflect.TypeOf((*func(accum, accum) accum)(nil)).Elem(), funcMakerAccumAccumГAccum)
	reflectx.RegisterFunc(reflect.TypeOf((*func(accum, typex.T) accum)(nil)).Elem(), funcMakerAccumTypex۰TГAccum)
	reflectx.RegisterFunc(reflect.TypeOf((*func(accum) []typex.T)(nil)).Elem(), funcMakerAccumГSliceOfTypex۰T)
	reflectx.RegisterFunc(reflect.TypeOf((*func())(nil)).Elem(), funcMakerГ)
	reflectx.RegisterFunc(reflect.TypeOf((*func() accum)(nil)).Elem(), funcMakerГAccum)
}

func wrapMakerCombineFn(fn interface{}) map[string]reflectx.Func {
	dfn := fn.(*combineFn)
	return map[string]reflectx.Func{
		"AddInput":          reflectx.MakeFunc(func(a0 accum, a1 typex.T) accum { return dfn.AddInput(a0, a1) }),
		"CreateAccumulator": reflectx.MakeFunc(func() accum { return dfn.CreateAccumulator() }),
		"ExtractOutput":     reflectx.MakeFunc(func(a0 accum) []typex.T { return dfn.ExtractOutput(a0) }),
		"MergeAccumulators": reflectx.MakeFunc(func(a0 accum, a1 accum) accum { return dfn.MergeAccumulators(a0, a1) }),
		"Setup":             reflectx.MakeFunc(func() { dfn.Setup() }),
	}
}

type callerAccumAccumГAccum struct {
	fn func(accum, accum) accum
}

func funcMakerAccumAccumГAccum(fn interface{}) reflectx.Func {
	f := fn.(func(accum, accum) accum)
	return &callerAccumAccumГAccum{fn: f}
}

func (c *callerAccumAccumГAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerAccumAccumГAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerAccumAccumГAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(accum), args[1].(accum))
	return []interface{}{out0}
}

func (c *callerAccumAccumГAccum) Call2x1(arg0, arg1 interface{}) interface{} {
	return c.fn(arg0.(accum), arg1.(accum))
}

type callerAccumTypex۰TГAccum struct {
	fn func(accum, typex.T) accum
}

func funcMakerAccumTypex۰TГAccum(fn interface{}) reflectx.Func {
	f := fn.(func(accum, typex.T) accum)
	return &callerAccumTypex۰TГAccum{fn: f}
}

func (c *callerAccumTypex۰TГAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerAccumTypex۰TГAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerAccumTypex۰TГAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(accum), args[1].(typex.T))
	return []interface{}{out0}
}

func (c *callerAccumTypex۰TГAccum) Call2x1(arg0, arg1 interface{}) interface{} {
	return c.fn(arg0.(accum), arg1.(typex.T))
}

type callerAccumГSliceOfTypex۰T struct {
	fn func(accum) []typex.T
}

func funcMakerAccumГSliceOfTypex۰T(fn interface{}) reflectx.Func {
	f := fn.(func(accum) []typex.T)
	return &callerAccumГSliceOfTypex۰T{fn: f}
}

func (c *callerAccumГSliceOfTypex۰T) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerAccumГSliceOfTypex۰T) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerAccumГSliceOfTypex۰T) Call(args []interface{}) []interface{} {
	out0 := c.fn(args[0].(accum))
	return []interface{}{out0}
}

func (c *callerAccumГSliceOfTypex۰T) Call1x1(arg0 interface{}) interface{} {
	return c.fn(arg0.(accum))
}

type callerГ struct {
	fn func()
}

func funcMakerГ(fn interface{}) reflectx.Func {
	f := fn.(func())
	return &callerГ{fn: f}
}

func (c *callerГ) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГ) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГ) Call(args []interface{}) []interface{} {
	c.fn()
	return []interface{}{}
}

func (c *callerГ) Call0x0() {
	c.fn()
}

type callerГAccum struct {
	fn func() accum
}

func funcMakerГAccum(fn interface{}) reflectx.Func {
	f := fn.(func() accum)
	return &callerГAccum{fn: f}
}

func (c *callerГAccum) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГAccum) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГAccum) Call(args []interface{}) []interface{} {
	out0 := c.fn()
	return []interface{}{out0}
}

func (c *callerГAccum) Call0x1() interface{} {
	return c.fn()
}

// DO NOT MODIFY: GENERATED CODE
