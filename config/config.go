package config

import (
	"github.com/simple-music/gateway/args"
	"github.com/simple-music/gateway/env"
)

var (
	ServiceHost = "0.0.0.0"
	ServicePort = args.GetString("service_port", "8080")

	TokenSigningKey = env.GetVar("TOKEN_SIGNING_KEY", "secret")

	DiscoveryHost = env.GetVar("SIMPLE_MUSIC_REGISTRY_HOST", "127.0.0.1")
	DiscoveryPort = env.GetVar("SIMPLE_MUSIC_REGISTRY_PORT", "8500")

	AuthServiceID       = env.GetVar("SIMPLE_MUSIC_AUTH_SERVICE_ID", "auth")
	AuthServicePassword = env.GetVar("SIMPLE_MUSIC_AUTH_SERVICE_PASSWORD", "secret")

	MusiciansServiceID       = env.GetVar("SIMPLE_MUSIC_MUSICIANS_SERVICE_ID", "musicians")
	MusiciansServicePassword = env.GetVar("SIMPLE_MUSIC_MUSICIANS_SERVICE_PASSWORD", "secret")

	SubscriptionsServiceID       = env.GetVar("SIMPLE_MUSIC_SUBSCRIPTIONS_SERVICE_ID", "subscriptions")
	SubscriptionsServicePassword = env.GetVar("SIMPLE_MUSIC_SUBSCRIPTIONS_SERVICE_PASSWORD", "secret")

	AvatarsServiceID       = env.GetVar("SIMPLE_MUSIC_AVATARS_SERVICE_ID", "avatars")
	AvatarsServicePassword = env.GetVar("SIMPLE_MUSIC_AVATARS_SERVICE_ID", "secret")
)
