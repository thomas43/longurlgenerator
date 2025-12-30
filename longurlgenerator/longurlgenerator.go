package longurlgenerator

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"strings"
	"time"
)

const (
	baseURL            = "https://reallylongurl.com"
	minimumQueryParams = 50
	maximumQueryParams = 72
)

func GenerateLongURL(inputURL string) (string, error) {
	// Parse the URL to validate it
	_, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	// Generate a unique identifier based on the URL hash
	uniqueID := generateUniqueIdentifier(inputURL)

	// Generate absurd query parameters
	params := generateAbsurdParams(inputURL)

	// Construct the long URL with our domain
	//https://reallylongurl.com/<hash>?<queryParams>
	longURL := fmt.Sprintf("%s/%s?%s", baseURL, uniqueID, params)

	return longURL, nil
}

func generateUniqueIdentifier(inputURL string) string {
	// Use SHA-256 to hash the URL
	hash := sha256.Sum256([]byte(inputURL))

	hashStr := hex.EncodeToString(hash[:])
	return hashStr
}

func generateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func generateUUID() string {
	return fmt.Sprintf("%s-%s-%s-%s-%s",
		generateRandomID(8),
		generateRandomID(4),
		generateRandomID(4),
		generateRandomID(4),
		generateRandomID(12))
}

func generateAbsurdParams(originalURL string) string {
	now := time.Now()
	encoded := base64.StdEncoding.EncodeToString([]byte(originalURL))
	hash := sha256.Sum256([]byte(originalURL))

	// Giant list of all possible absurd parameters
	allParams := []string{
		// Session and tracking IDs
		fmt.Sprintf("session_id=%s", generateRandomID(32)),
		fmt.Sprintf("request_id=%s", generateRandomID(16)),
		fmt.Sprintf("correlation_id=%s", generateUUID()),
		fmt.Sprintf("tracking_pixel_id=%s", generateRandomID(24)),
		fmt.Sprintf("transaction_id=%s", generateUUID()),
		fmt.Sprintf("visitor_id=%s", generateRandomID(28)),
		fmt.Sprintf("client_id=%s", generateRandomID(20)),
		fmt.Sprintf("anonymous_id=%s", generateUUID()),
		fmt.Sprintf("user_id=%s", generateUUID()),
		fmt.Sprintf("account_id=%s", generateRandomID(18)),
		fmt.Sprintf("organization_id=%s", generateUUID()),
		fmt.Sprintf("workspace_id=%s", generateRandomID(22)),
		fmt.Sprintf("team_id=%s", generateRandomID(16)),
		fmt.Sprintf("project_id=%s", generateUUID()),
		fmt.Sprintf("trace_id=%s", generateRandomID(32)),
		fmt.Sprintf("span_id=%s", generateRandomID(16)),
		fmt.Sprintf("parent_span_id=%s", generateRandomID(16)),
		fmt.Sprintf("conversation_id=%s", generateUUID()),
		fmt.Sprintf("interaction_id=%s", generateRandomID(24)),
		fmt.Sprintf("event_id=%s", generateUUID()),

		// Timestamps in various formats
		fmt.Sprintf("timestamp_unix=%d", now.Unix()),
		fmt.Sprintf("timestamp_iso=%s", now.Format(time.RFC3339)),
		fmt.Sprintf("timestamp_epoch_ms=%d", now.UnixMilli()),
		fmt.Sprintf("timestamp_utc=%s", now.UTC().Format(time.RFC1123)),
		fmt.Sprintf("timestamp_rfc2822=%s", now.Format(time.RFC1123Z)),
		fmt.Sprintf("server_time=%d", now.Unix()),
		fmt.Sprintf("client_time=%d", now.UnixMilli()),
		fmt.Sprintf("request_timestamp=%s", now.Format("2006-01-02T15:04:05Z07:00")),
		fmt.Sprintf("created_at=%d", now.Unix()),
		fmt.Sprintf("updated_at=%d", now.Unix()),
		fmt.Sprintf("processed_at=%d", now.UnixMilli()),
		fmt.Sprintf("received_at=%s", now.Format(time.RFC3339Nano)),

		// UTM and marketing parameters
		"utm_source=organic-social-media-platform-referral",
		"utm_medium=referral-link-sharing-mechanism-protocol",
		"utm_campaign=user-generated-content-distribution-initiative",
		"utm_term=highly-relevant-search-keywords-semantic-analysis",
		"utm_content=premium-quality-engagement-content-variant-a",
		"utm_id=campaign-identifier-unique-tracking-string",
		"utm_source_platform=cross-platform-social-aggregator-service",
		"utm_creative_format=responsive-display-advertisement-unit",
		"utm_marketing_tactic=inbound-content-marketing-strategy",
		"marketing_channel=omnichannel-digital-experience-touchpoint",
		"campaign_name=quarterly-engagement-growth-initiative-q4",
		"campaign_type=awareness-consideration-conversion-funnel",
		"ad_group=targeted-demographic-psychographic-segment",
		"ad_group_id=advertisement-group-unique-identifier",
		"creative_id=creative-asset-version-identifier",
		"placement_id=advertising-placement-location-code",
		"keyword=search-engine-marketing-keyword-match",
		"match_type=broad-phrase-exact-keyword-matching",
		"ad_position=search-engine-results-page-position",
		"network=search-display-video-network-type",

		// Encoded data and hashes
		fmt.Sprintf("original_url_encoded=%s", encoded),
		fmt.Sprintf("url_hash=%x", hash),
		fmt.Sprintf("payload=%s", base64.StdEncoding.EncodeToString([]byte("meaningless-payload-data"))),
		fmt.Sprintf("signature=%s", generateRandomID(64)),
		fmt.Sprintf("checksum=%s", generateRandomID(32)),
		fmt.Sprintf("hmac=%s", generateRandomID(64)),
		fmt.Sprintf("digest=%s", generateRandomID(40)),
		fmt.Sprintf("fingerprint=%s", generateRandomID(48)),
		fmt.Sprintf("verification_token=%s", generateRandomID(32)),
		fmt.Sprintf("integrity_hash=%s", generateRandomID(56)),

		// Device and browser information
		fmt.Sprintf("device_fingerprint=%s", generateRandomID(40)),
		fmt.Sprintf("canvas_fingerprint=%s", generateRandomID(32)),
		fmt.Sprintf("audio_fingerprint=%s", generateRandomID(28)),
		fmt.Sprintf("webgl_fingerprint=%s", generateRandomID(36)),
		"browser_session_storage_available=true-verified",
		"browser_local_storage_available=true-verified",
		"browser_indexed_db_available=true-verified",
		"cookies_enabled=definitely-yes-absolutely-confirmed",
		"third_party_cookies_enabled=deprecated-browser-policy",
		"screen_resolution=1920x1080-high-definition-display",
		"screen_width=1920-pixels-horizontal",
		"screen_height=1080-pixels-vertical",
		"available_screen_width=1920-pixels-available",
		"available_screen_height=1040-pixels-available",
		"color_depth=24-bit-true-color-display",
		"pixel_ratio=2.0-retina-display-hdpi",
		"viewport_width=1920-pixels-initial-viewport",
		"viewport_height=1080-pixels-initial-viewport",
		"device_memory=8gb-ram-available-javascript",
		"hardware_concurrency=8-logical-processors-available",
		"max_touch_points=0-no-touch-support",
		"platform=win32-windows-operating-system-platform",
		"user_agent=mozilla-5.0-chrome-compatible-browser",
		"browser_name=chromium-based-web-browser-engine",
		"browser_version=120.0.6099.129-stable-release",
		"browser_major_version=120-current-major",
		"engine_name=blink-rendering-engine-webkit-fork",
		"engine_version=120.0.6099.129-engine-version",
		"os_name=windows-nt-10.0-latest-version",
		"os_version=10.0.19045-build-number",
		"device_type=desktop-workstation-computer-form-factor",
		"device_vendor=generic-pc-compatible-manufacturer",
		"device_model=desktop-standard-configuration",
		"cpu_architecture=x86-64-amd64-intel-compatible",
		"gpu_vendor=nvidia-amd-intel-graphics",
		"gpu_renderer=angle-direct3d11-opengl-backend",
		"touch_support=enabled-multitouch-gestures-available",
		"pointer_type=mouse-primary-input-device",
		"orientation=landscape-primary-screen-orientation",
		"vendor=google-inc-chromium-project",
		"vendor_sub=official-build-distribution",
		"product=gecko-compatibility-mode",
		"product_sub=20030107-gecko-version",

		// Location and internationalization
		"timezone_offset=-480-minutes-pst-pacific",
		"timezone=America-Los_Angeles-tz-database",
		"language_preference=en-US-english-united-states-locale",
		"language=en-primary-language-code",
		"languages=en-US,en;q=0.9-accepted-languages",
		"locale=en-US-american-english-locale-identifier",
		"country_code=US-united-states-america-iso",
		"country_name=United-States-of-America-full",
		"region=california-west-coast-state",
		"region_code=CA-california-abbreviation",
		"city=san-francisco-bay-area-location",
		"postal_code=94102-geographic-zip-code",
		"latitude=37.7749-degrees-north-coordinate",
		"longitude=-122.4194-degrees-west-coordinate",
		"accuracy=50-meters-geolocation-precision",
		"altitude=52-meters-above-sea-level",
		"altitude_accuracy=10-meters-altitude-precision",
		"heading=null-direction-not-available",
		"speed=null-velocity-not-available",
		"ip_address_hash=anonymized-for-privacy-gdpr-compliant",
		"ip_version=ipv4-internet-protocol-version",
		"isp=internet-service-provider-name-redacted",
		"asn=autonomous-system-number-redacted",
		"connection_type=wifi-wireless-broadband-connection",

		// Privacy and compliance
		"do_not_track=1-respectfully-ignored-by-default",
		"global_privacy_control=0-gpc-signal-not-set",
		"gdpr_consent=CP1234567890-tcf-v2-consent-string",
		"gdpr_applies=true-european-union-visitor",
		"ccpa_opt_out=california-privacy-rights-do-not-sell",
		"ccpa_applies=true-california-resident-indicator",
		"cookie_consent=all-categories-accepted-timestamp",
		"cookie_consent_necessary=true-required-cookies",
		"cookie_consent_functional=true-functional-cookies",
		"cookie_consent_analytics=true-analytics-cookies",
		"cookie_consent_advertising=true-advertising-cookies",
		"privacy_policy_version=2024.1.0-latest-revision",
		"privacy_policy_accepted=true-user-acknowledged",
		"terms_of_service_version=2024.1.0-current",
		"terms_accepted=true-user-agreement-confirmed",
		"age_verification=confirmed-over-18-years-old",
		"age_gate_passed=true-minimum-age-requirement",
		"data_processing_consent=granted-explicitly-informed",
		"marketing_consent=granted-opt-in-confirmed",
		"personalization_consent=granted-tailored-experience",
		"data_retention_acknowledged=true-policy-understood",

		// Performance and resource timing
		"performance_timing=enabled-for-analytics-collection",
		"navigation_timing=navigation-timing-api-level-2",
		"resource_timing=resource-timing-api-enabled",
		"paint_timing=first-contentful-paint-metric",
		"largest_contentful_paint=2847-milliseconds-lcp",
		"first_input_delay=12-milliseconds-fid-metric",
		"cumulative_layout_shift=0.045-cls-score",
		"time_to_interactive=3240-milliseconds-tti",
		"total_blocking_time=287-milliseconds-tbt",
		"speed_index=2456-milliseconds-si-metric",
		"dom_content_loaded=1847-milliseconds-dcl",
		"load_complete=3456-milliseconds-onload",
		"first_byte=234-milliseconds-ttfb-server",
		"dns_lookup=45-milliseconds-dns-resolution",
		"tcp_connection=89-milliseconds-tcp-handshake",
		"tls_negotiation=123-milliseconds-ssl-handshake",
		"request_time=178-milliseconds-http-request",
		"response_time=234-milliseconds-http-response",
		"dom_processing=892-milliseconds-dom-parse",
		"render_time=445-milliseconds-rendering",

		// Features and capabilities
		"javascript_enabled=true-scripts-allowed-executed",
		"webgl_support=enabled-3d-graphics-hardware-acceleration",
		"webgl_version=2.0-opengl-es-3.0-context",
		"webgl2_support=true-modern-graphics-api",
		"web_audio_api=supported-audio-context-available",
		"web_rtc=supported-real-time-communication",
		"web_workers=supported-background-threads",
		"service_workers=supported-progressive-web-app",
		"push_notifications=supported-browser-capability",
		"notifications_permission=default-not-granted",
		"geolocation_permission=prompt-not-decided",
		"camera_permission=denied-user-declined",
		"microphone_permission=denied-user-declined",
		"midi_support=true-web-midi-api",
		"payment_request=supported-web-payments",
		"credential_management=supported-credentials-api",
		"web_assembly=supported-wasm-execution",
		"shared_array_buffer=supported-cross-origin-isolated",
		"web_usb=supported-usb-device-access",
		"web_bluetooth=not-supported-browser-limitation",
		"nfc=not-supported-near-field-communication",
		"local_fonts=supported-font-enumeration-api",

		// Referrer and navigation
		"referrer_type=direct-navigation-user-intent-typed",
		"referrer_url=none-direct-traffic-source",
		"referrer_domain=direct-no-referring-domain",
		"landing_page=current-entry-point-url-path",
		"entry_page=initial-session-landing-page",
		"exit_page=null-session-still-active",
		"previous_page=document-referrer-same-origin",
		"page_depth=5-clicks-from-home-page",
		"visit_count=42-lifetime-visits-returning",
		"session_count=18-total-sessions-historical",
		"return_visitor=true-recognized-user-profile",
		"new_visitor=false-existing-profile-identified",
		"days_since_last_visit=7-days-elapsed",
		"visits_today=2-sessions-current-day",
		"bounce_rate=calculated-engagement-metric-percentage",
		"engagement_score=87-calculated-user-quality",

		// A/B testing and experiments
		"ab_test_variant=control-group-baseline-original",
		"ab_test_id=test-unique-identifier-uuid",
		"experiment_id=feature-flag-test-123-active",
		"experiment_name=homepage-hero-test-variant",
		"variant_id=treatment-arm-b-experimental",
		"variant_name=new-design-version-two",
		"cohort=user-segment-alpha-targeted",
		"cohort_id=segment-identifier-uuid",
		"feature_flag_bundle=enabled-features-list-json",
		"feature_flags=feature-a:true,feature-b:false",
		"optimization_group=conversion-rate-test-checkout",
		"personalization_id=tailored-experience-identifier",
		"recommendation_engine=collaborative-filtering-v2",
		"test_start_date=2024-01-15-experiment-launch",
		"test_allocation=random-50-50-split-traffic",

		// Session information
		"session_duration=1847-seconds-elapsed-time",
		"session_start=2024-01-15T14:23:45Z-iso-timestamp",
		"page_views=12-pages-this-session-count",
		"page_views_total=247-lifetime-pageviews",
		"events_fired=37-interactions-tracked-session",
		"scroll_depth=85-percent-scrolled-maximum",
		"scroll_depth_px=4280-pixels-scrolled-vertical",
		"time_on_page=247-seconds-engaged-current",
		"active_time=189-seconds-active-engagement",
		"idle_time=58-seconds-inactive-idle",
		"interaction_count=14-clicks-recorded-session",
		"click_count=9-mouse-clicks-tracked",
		"keypress_count=145-keyboard-inputs-recorded",
		"form_submissions=2-forms-completed-submitted",
		"form_starts=3-forms-initiated-engagement",
		"video_plays=1-video-interactions-playback",
		"video_completion=0-videos-watched-completely",
		"downloads=0-file-downloads-initiated",
		"outbound_clicks=2-external-links-clicked",
		"social_shares=0-social-sharing-interactions",

		// Network and connection
		"connection_type=wifi-wireless-broadband-network",
		"connection_speed=fast-4g-lte-mobile-data",
		"effective_connection_type=4g-network-quality",
		"network_quality=excellent-low-latency-stable",
		"bandwidth_estimate=50-mbps-download-speed",
		"downlink=10-mbps-effective-bandwidth",
		"rtt=50-milliseconds-round-trip-time",
		"save_data=false-data-saver-mode-disabled",
		"protocol=https-secure-connection-tls",
		"protocol_version=http-2-multiplexed-streams",
		"tls_version=1.3-modern-encryption-standard",
		"cipher_suite=TLS_AES_128_GCM_SHA256-cipher",
		"certificate_issuer=lets-encrypt-authority-x3",
		"certificate_valid_until=2025-06-15-expiry",
		"ocsp_stapling=enabled-certificate-validation",
		"hsts=enforced-strict-transport-security",
		"http_strict_transport_security=max-age-31536000",

		// Random security tokens
		fmt.Sprintf("nonce=%s", generateRandomID(16)),
		fmt.Sprintf("salt=%s", generateRandomID(24)),
		fmt.Sprintf("token=%s", generateRandomID(48)),
		fmt.Sprintf("api_key=%s", generateRandomID(32)),
		fmt.Sprintf("auth_token=%s", generateRandomID(40)),
		fmt.Sprintf("bearer_token=%s", generateRandomID(64)),
		fmt.Sprintf("refresh_token=%s", generateRandomID(48)),
		fmt.Sprintf("access_token=%s", generateRandomID(52)),
		fmt.Sprintf("csrf_token=%s", generateRandomID(32)),
		fmt.Sprintf("state_token=%s", generateRandomID(24)),
		fmt.Sprintf("oauth_token=%s", generateRandomID(40)),
		fmt.Sprintf("oauth_verifier=%s", generateRandomID(32)),
		fmt.Sprintf("code_verifier=%s", generateRandomID(128)),
		fmt.Sprintf("code_challenge=%s", generateRandomID(64)),

		// Completely absurd tech buzzwords
		"quantum_entanglement_id=superposition-state-collapsed-measured",
		"blockchain_hash=distributed-ledger-proof-of-work-sha256",
		"smart_contract_address=0x742d35Cc6634C0532925a3b844Bc9e7595f0bEb",
		"nft_token_id=erc-721-non-fungible-token-identifier",
		"web3_wallet_address=0x71C7656EC7ab88b098defB751B7401B5f6d8976F",
		"metaverse_coordinates=x:1247,y:8934,z:42-virtual-location",
		"ai_confidence_score=0.9847-machine-learning-prediction",
		"neural_network_layer=deep-learning-inference-layer-42",
		"ml_model_version=v2.4.1-trained-2024-01-15",
		"inference_latency=23-milliseconds-model-execution",
		"training_dataset=proprietary-supervised-learning-corpus",
		"cloud_region=us-west-2a-availability-zone-primary",
		"cloud_provider=aws-amazon-web-services-ec2",
		"container_id=kubernetes-pod-identifier-docker",
		"pod_name=frontend-deployment-7d4f9c8b6-xk2lp",
		"namespace=production-default-kubernetes-namespace",
		"cluster_name=prod-us-west-2-eks-cluster",
		"microservice_name=user-authentication-service-v2",
		"microservice_trace=distributed-tracing-id-jaeger",
		"service_mesh=istio-sidecar-proxy-envoy",
		"load_balancer_node=round-robin-selection-algorithm",
		"backend_server=server-42-load-balanced-pool",
		"cache_hit_ratio=87-percent-cached-redis",
		"cache_key=user:1234:profile-memcached-key",
		"cdn_pop=edge-location-nearest-cloudflare",
		"cdn_cache_status=HIT-served-from-edge",
		"edge_location=SFO-san-francisco-pop",
		"ssl_cipher_suite=ECDHE-RSA-AES256-GCM-SHA384",
		"http2_stream_id=15-multiplexed-connection",
		"http3_quic_version=draft-29-udp-protocol",
		"websocket_protocol=wss-bidirectional-communication",
		"graphql_operation=query-user-profile-batched",
		"rest_api_version=v2-restful-web-services",
		"grpc_method=GetUser-remote-procedure-call",
		"message_queue=rabbitmq-amqp-broker-exchange",
		"event_bus=kafka-topic-user-events-partition-3",
		"database_shard=shard-7-horizontal-partitioning",
		"replica_set=mongodb-primary-replica-set",
		"read_preference=secondary-preferred-load-distribution",
		"write_concern=majority-acknowledged-durability",
		"transaction_isolation=read-committed-acid-level",
		"service_worker_version=v24-progressive-web-app",
		"manifest_version=web-app-manifest-v2-json",
		"workbox_version=6.5.4-service-worker-library",
		"pwa_install_prompt=deferred-add-to-homescreen",
		"notification_permission=granted-push-enabled",
		"background_sync=registered-offline-capability",
		"indexed_db_version=3-local-storage-database",
	}

	// Randomly select between min/max parameters
	numParams := rand.Intn(maximumQueryParams-minimumQueryParams) + minimumQueryParams

	// Shuffle the parameters
	rand.Shuffle(len(allParams), func(i, j int) {
		allParams[i], allParams[j] = allParams[j], allParams[i]
	})

	// Take the first numParams
	if numParams > len(allParams) {
		numParams = len(allParams)
	}
	selectedParams := allParams[:numParams]

	return strings.Join(selectedParams, "&")
}
