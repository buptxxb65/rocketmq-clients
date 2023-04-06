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

package main

import (
	"log"
	"os"
	"time"

	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
)

const (
	Topic         = "xxxxxx"
	ConsumerGroup = "xxxxxx"
	Endpoint      = "xxxxxx"
	AccessKey     = "xxxxxx"
	SecretKey     = "xxxxxx"
)

var (
	// maximum waiting time for receive func
	awaitDuration = time.Second * 5
	// invisibleDuration should > 20s
	invisibleDuration = time.Second * 20
	// messageViewCacheSize the maximum received message in the cache
	messageViewCacheSize = 10
	// receive messages in a loop
)

func main() {
	// log to console
	os.Setenv("mq.consoleAppender.enabled", "true")
	rmq_client.ResetLogger()
	// new pushConsumer instance
	pushConsumer, err := rmq_client.NewPushConsumer(&rmq_client.Config{
		Endpoint:      Endpoint,
		ConsumerGroup: ConsumerGroup,
		Credentials: &credentials.SessionCredentials{
			AccessKey:    AccessKey,
			AccessSecret: SecretKey,
		},
	},
		&rmq_client.FuncPushConsumerOption{
			FuncConsumerOption: rmq_client.WithAwaitDuration(awaitDuration),
		},
		&rmq_client.FuncPushConsumerOption{
			FuncConsumerOption: rmq_client.WithSubscriptionExpressions(map[string]*rmq_client.FilterExpression{
				Topic: rmq_client.SUB_ALL,
			}),
		},
		rmq_client.WithInvisibleDuration(invisibleDuration),
		rmq_client.WithMessageViewCacheSize(messageViewCacheSize),
	)
	if err != nil {
		log.Fatal(err)
	}
	// start simpleConsumer
	err = pushConsumer.Start()
	if err != nil {
		log.Fatal(err)
	}
	// graceful stop simpleConsumer
	defer pushConsumer.GracefulStop()

	// run for a while
	time.Sleep(time.Minute)
}
