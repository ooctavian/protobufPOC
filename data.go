package main

import (
	"time"

	"protobufPOC/schema"

	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	boolTrue  = true
	boolFalse = false

	articles = &schema.Articles{
		Data: []*schema.Article{
			{
				Id:    1,
				Title: "The C Programming Language",
				Authors: []*schema.Author{
					{
						FirstName: "Dennis",
						LastName:  "Ritchie",
					},
					{
						FirstName: "Brian",
						LastName:  "Kernighan",
					}},
				Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Categories: []uint32{
					uint32(3),
					uint32(6),
				},
				Draft:        &boolTrue,
				DateCreated:  timestamppb.New(time.Now()),
				DateModefied: timestamppb.Now(),
			},
			{
				Id:    2,
				Title: "The Rust Programming Language",
				Authors: []*schema.Author{
					{
						FirstName: "Carol",
						LastName:  "Nichols",
					},
					{
						FirstName: "Steve",
						LastName:  "Klabnik",
					},
				},
				Content: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Draft:   &boolFalse,
				Categories: []uint32{
					uint32(3),
					uint32(2),
					uint32(6),
				},
				DateCreated:  timestamppb.New(time.Now()),
				DateModefied: timestamppb.Now(),
			},
			{
				Id:    3,
				Title: "Programming PHP",
				Authors: []*schema.Author{
					{
						FirstName: "Rasmus",
						LastName:  "Leordfor",
					},
				},
				Categories: []uint32{
					uint32(3),
					uint32(2),
				},
				Content:      "PHP is a simple yet powerful open-source scripting language for creating dynamic web content. The millions of web sites powered by PHP are testament to its popularity and ease of use",
				Draft:        &boolTrue,
				DateCreated:  timestamppb.New(time.Now()),
				DateModefied: timestamppb.Now(),
			},
			{
				Id:    4,
				Title: "Programming: Principles and Practice Using C++",
				Authors: []*schema.Author{
					{
						FirstName: "Bjarne",
						LastName:  "Stroustrup",
					},
				},
				Categories: []uint32{
					uint32(3),
				},
				Content:      "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
				Draft:        &boolFalse,
				DateCreated:  timestamppb.New(time.Now()),
				DateModefied: timestamppb.Now(),
			},
		},
	}
)
