/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package golang

import "time"

type FilterExpressionType int32

const (
	SQL92 FilterExpressionType = iota
	TAG
	UNSPECIFIED
)

type FilterExpression struct {
	expression     string
	expressionType FilterExpressionType
}

var SUB_ALL = NewFilterExpression("*")

var NewFilterExpression = func(expression string) *FilterExpression {
	return &FilterExpression{
		expression:     expression,
		expressionType: TAG,
	}
}

var NewFilterExpressionWithType = func(expression string, expressionType FilterExpressionType) *FilterExpression {
	return &FilterExpression{
		expression:     expression,
		expressionType: expressionType,
	}
}

type consumerOptions struct {
	subscriptionExpressions map[string]*FilterExpression
	awaitDuration           time.Duration
	clientFunc              NewClientFunc
}

var defaultConsumerOptions = consumerOptions{
	clientFunc: NewClient,
}

// A ConsumerOption sets options such as tag, etc.
type ConsumerOption interface {
	apply(*consumerOptions)
}

// FuncConsumerOption wraps a function that modifies options into an implementation of
// the Option interface.
type FuncConsumerOption struct {
	f func(*consumerOptions)
}

func (fo *FuncConsumerOption) apply(do *consumerOptions) {
	fo.f(do)
}

func newFuncConsumerOption(f func(*consumerOptions)) *FuncConsumerOption {
	return &FuncConsumerOption{
		f: f,
	}
}

func WithSubscriptionExpressions(subscriptionExpressions map[string]*FilterExpression) *FuncConsumerOption {
	return newFuncConsumerOption(func(o *consumerOptions) {
		o.subscriptionExpressions = subscriptionExpressions
	})
}

func WithAwaitDuration(awaitDuration time.Duration) *FuncConsumerOption {
	return newFuncConsumerOption(func(o *consumerOptions) {
		o.awaitDuration = awaitDuration
	})
}
