// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.15.2
// source: vito-stt-client.proto

package stt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Indicates the type of speech event.
type DecoderResponse_SpeechEventType int32

const (
	// No speech event specified.
	DecoderResponse_SPEECH_EVENT_UNSPECIFIED DecoderResponse_SpeechEventType = 0
	// This event indicates that the server has detected the end of the user's
	// speech utterance and expects no additional speech. Therefore, the server
	// will not process additional audio (although it may subsequently return
	// additional results). The client should stop sending additional audio
	// data, half-close the gRPC connection, and wait for any additional results
	// until the server closes the gRPC connection. This event is only sent if
	// `single_utterance` was set to `true`, and is not used otherwise.
	DecoderResponse_END_OF_SINGLE_UTTERANCE DecoderResponse_SpeechEventType = 1
)

// Enum value maps for DecoderResponse_SpeechEventType.
var (
	DecoderResponse_SpeechEventType_name = map[int32]string{
		0: "SPEECH_EVENT_UNSPECIFIED",
		1: "END_OF_SINGLE_UTTERANCE",
	}
	DecoderResponse_SpeechEventType_value = map[string]int32{
		"SPEECH_EVENT_UNSPECIFIED": 0,
		"END_OF_SINGLE_UTTERANCE":  1,
	}
)

func (x DecoderResponse_SpeechEventType) Enum() *DecoderResponse_SpeechEventType {
	p := new(DecoderResponse_SpeechEventType)
	*p = x
	return p
}

func (x DecoderResponse_SpeechEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DecoderResponse_SpeechEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_vito_stt_client_proto_enumTypes[0].Descriptor()
}

func (DecoderResponse_SpeechEventType) Type() protoreflect.EnumType {
	return &file_vito_stt_client_proto_enumTypes[0]
}

func (x DecoderResponse_SpeechEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DecoderResponse_SpeechEventType.Descriptor instead.
func (DecoderResponse_SpeechEventType) EnumDescriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{1, 0}
}

type DecoderConfig_AudioEncoding int32

const (
	// Not specified.
	DecoderConfig_ENCODING_UNSPECIFIED DecoderConfig_AudioEncoding = 0
	// Uncompressed 16-bit signed little-endian samples (Linear PCM).
	DecoderConfig_LINEAR16 DecoderConfig_AudioEncoding = 1
	DecoderConfig_WAV      DecoderConfig_AudioEncoding = 2
	// `FLAC` (Free Lossless Audio
	// Codec) is the recommended encoding because it is
	// lossless--therefore recognition is not compromised--and
	// requires only about half the bandwidth of `LINEAR16`. `FLAC` stream
	// encoding supports 16-bit and 24-bit samples, however, not all fields in
	// `STREAMINFO` are supported.
	DecoderConfig_FLAC DecoderConfig_AudioEncoding = 3
	// 8-bit samples that compand 14-bit audio samples using G.711 PCMU/mu-law.
	DecoderConfig_MULAW DecoderConfig_AudioEncoding = 4
	DecoderConfig_ALAW  DecoderConfig_AudioEncoding = 5
	// Adaptive Multi-Rate Narrowband codec. `sample_rate_hertz` must be 8000.
	DecoderConfig_AMR DecoderConfig_AudioEncoding = 6
	// Adaptive Multi-Rate Wideband codec. `sample_rate_hertz` must be 16000.
	DecoderConfig_AMR_WB DecoderConfig_AudioEncoding = 7
	// Opus encoded audio frames in Ogg container
	// ([OggOpus](https://wiki.xiph.org/OggOpus)).
	// `sample_rate_hertz` must be one of 8000, 12000, 16000, 24000, or 48000.
	DecoderConfig_OGG_OPUS DecoderConfig_AudioEncoding = 8
	// Opus encoded audio frames without container
	DecoderConfig_OPUS DecoderConfig_AudioEncoding = 9
)

// Enum value maps for DecoderConfig_AudioEncoding.
var (
	DecoderConfig_AudioEncoding_name = map[int32]string{
		0: "ENCODING_UNSPECIFIED",
		1: "LINEAR16",
		2: "WAV",
		3: "FLAC",
		4: "MULAW",
		5: "ALAW",
		6: "AMR",
		7: "AMR_WB",
		8: "OGG_OPUS",
		9: "OPUS",
	}
	DecoderConfig_AudioEncoding_value = map[string]int32{
		"ENCODING_UNSPECIFIED": 0,
		"LINEAR16":             1,
		"WAV":                  2,
		"FLAC":                 3,
		"MULAW":                4,
		"ALAW":                 5,
		"AMR":                  6,
		"AMR_WB":               7,
		"OGG_OPUS":             8,
		"OPUS":                 9,
	}
)

func (x DecoderConfig_AudioEncoding) Enum() *DecoderConfig_AudioEncoding {
	p := new(DecoderConfig_AudioEncoding)
	*p = x
	return p
}

func (x DecoderConfig_AudioEncoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DecoderConfig_AudioEncoding) Descriptor() protoreflect.EnumDescriptor {
	return file_vito_stt_client_proto_enumTypes[1].Descriptor()
}

func (DecoderConfig_AudioEncoding) Type() protoreflect.EnumType {
	return &file_vito_stt_client_proto_enumTypes[1]
}

func (x DecoderConfig_AudioEncoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DecoderConfig_AudioEncoding.Descriptor instead.
func (DecoderConfig_AudioEncoding) EnumDescriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{5, 0}
}

// The request message containing the user's name and how many greetings
// they want.
type DecoderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The streaming request, which is either a streaming config or audio content.
	//
	// Types that are assignable to StreamingRequest:
	//	*DecoderRequest_StreamingConfig
	//	*DecoderRequest_AudioContent
	StreamingRequest isDecoderRequest_StreamingRequest `protobuf_oneof:"streaming_request"`
}

func (x *DecoderRequest) Reset() {
	*x = DecoderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecoderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecoderRequest) ProtoMessage() {}

func (x *DecoderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecoderRequest.ProtoReflect.Descriptor instead.
func (*DecoderRequest) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{0}
}

func (m *DecoderRequest) GetStreamingRequest() isDecoderRequest_StreamingRequest {
	if m != nil {
		return m.StreamingRequest
	}
	return nil
}

func (x *DecoderRequest) GetStreamingConfig() *DecoderConfig {
	if x, ok := x.GetStreamingRequest().(*DecoderRequest_StreamingConfig); ok {
		return x.StreamingConfig
	}
	return nil
}

func (x *DecoderRequest) GetAudioContent() []byte {
	if x, ok := x.GetStreamingRequest().(*DecoderRequest_AudioContent); ok {
		return x.AudioContent
	}
	return nil
}

type isDecoderRequest_StreamingRequest interface {
	isDecoderRequest_StreamingRequest()
}

type DecoderRequest_StreamingConfig struct {
	// Provides information to the recognizer that specifies how to process the
	// request. The first `StreamingRecognizeRequest` message must contain a
	// `streaming_config`  message.
	StreamingConfig *DecoderConfig `protobuf:"bytes,1,opt,name=streaming_config,json=streamingConfig,proto3,oneof"`
}

type DecoderRequest_AudioContent struct {
	// The audio data to be recognized. Sequential chunks of audio data are sent
	// in sequential `StreamingRecognizeRequest` messages. The first
	// `StreamingRecognizeRequest` message must not contain `audio_content` data
	// and all subsequent `StreamingRecognizeRequest` messages must contain
	// `audio_content` data. The audio bytes must be encoded as specified in
	// `RecognitionConfig`. Note: as with all bytes fields, proto buffers use a
	// pure binary representation (not base64). See
	// [content limits](https://cloud.google.com/speech-to-text/quotas#content).
	AudioContent []byte `protobuf:"bytes,2,opt,name=audio_content,json=audioContent,proto3,oneof"`
}

func (*DecoderRequest_StreamingConfig) isDecoderRequest_StreamingRequest() {}

func (*DecoderRequest_AudioContent) isDecoderRequest_StreamingRequest() {}

// A response message containing a greeting
type DecoderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error bool `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	// This repeated list contains zero or more results that
	// correspond to consecutive portions of the audio currently being processed.
	// It contains zero or one `is_final=true` result (the newly settled portion),
	// followed by zero or more `is_final=false` results (the interim results).
	Results []*StreamingRecognitionResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
	// Indicates the type of speech event.
	SpeechEventType DecoderResponse_SpeechEventType `protobuf:"varint,4,opt,name=speech_event_type,json=speechEventType,proto3,enum=online_decoder.DecoderResponse_SpeechEventType" json:"speech_event_type,omitempty"`
}

func (x *DecoderResponse) Reset() {
	*x = DecoderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecoderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecoderResponse) ProtoMessage() {}

func (x *DecoderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecoderResponse.ProtoReflect.Descriptor instead.
func (*DecoderResponse) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{1}
}

func (x *DecoderResponse) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *DecoderResponse) GetResults() []*StreamingRecognitionResult {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *DecoderResponse) GetSpeechEventType() DecoderResponse_SpeechEventType {
	if x != nil {
		return x.SpeechEventType
	}
	return DecoderResponse_SPEECH_EVENT_UNSPECIFIED
}

// A streaming speech recognition result corresponding to a portion of the audio
// that is currently being processed.
type StreamingRecognitionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// May contain one or more recognition hypotheses (up to the
	// maximum specified in `max_alternatives`).
	// These alternatives are ordered in terms of accuracy, with the top (first)
	// alternative being the most probable, as ranked by the recognizer.
	Alternatives []*SpeechRecognitionAlternative `protobuf:"bytes,1,rep,name=alternatives,proto3" json:"alternatives,omitempty"`
	// If `false`, this `StreamingRecognitionResult` represents an
	// interim result that may change. If `true`, this is the final time the
	// speech service will return this particular `StreamingRecognitionResult`,
	// the recognizer will not return any further hypotheses for this portion of
	// the transcript and corresponding audio.
	IsFinal bool `protobuf:"varint,2,opt,name=is_final,json=isFinal,proto3" json:"is_final,omitempty"`
	// An estimate of the likelihood that the recognizer will not
	// change its guess about this interim result. Values range from 0.0
	// (completely unstable) to 1.0 (completely stable).
	// This field is only provided for interim results (`is_final=false`).
	// The default of 0.0 is a sentinel value indicating `stability` was not set.
	Stability float32 `protobuf:"fixed32,3,opt,name=stability,proto3" json:"stability,omitempty"`
	// duration of the audio.
	Duration int32 `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time offset of the start of this result relative to the
	// beginning of the audio.
	StartAt int32 `protobuf:"varint,5,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
}

func (x *StreamingRecognitionResult) Reset() {
	*x = StreamingRecognitionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamingRecognitionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamingRecognitionResult) ProtoMessage() {}

func (x *StreamingRecognitionResult) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamingRecognitionResult.ProtoReflect.Descriptor instead.
func (*StreamingRecognitionResult) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{2}
}

func (x *StreamingRecognitionResult) GetAlternatives() []*SpeechRecognitionAlternative {
	if x != nil {
		return x.Alternatives
	}
	return nil
}

func (x *StreamingRecognitionResult) GetIsFinal() bool {
	if x != nil {
		return x.IsFinal
	}
	return false
}

func (x *StreamingRecognitionResult) GetStability() float32 {
	if x != nil {
		return x.Stability
	}
	return 0
}

func (x *StreamingRecognitionResult) GetDuration() int32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *StreamingRecognitionResult) GetStartAt() int32 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

// Alternative hypotheses (a.k.a. n-best list).
type SpeechRecognitionAlternative struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Transcript text representing the words that the user spoke.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// The confidence estimate between 0.0 and 1.0. A higher number
	// indicates an estimated greater likelihood that the recognized words are
	// correct. This field is set only for the top alternative of a non-streaming
	// result or, of a streaming result where `is_final=true`.
	// This field is not guaranteed to be accurate and users should not rely on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was not set.
	Confidence float32 `protobuf:"fixed32,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// A list of word-specific information for each recognized word.
	// Note: When `enable_speaker_diarization` is true, you will see all the words
	// from the beginning of the audio.
	Words []*WordInfo `protobuf:"bytes,3,rep,name=words,proto3" json:"words,omitempty"` // Word-specific information for recognized words.
}

func (x *SpeechRecognitionAlternative) Reset() {
	*x = SpeechRecognitionAlternative{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechRecognitionAlternative) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechRecognitionAlternative) ProtoMessage() {}

func (x *SpeechRecognitionAlternative) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechRecognitionAlternative.ProtoReflect.Descriptor instead.
func (*SpeechRecognitionAlternative) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{3}
}

func (x *SpeechRecognitionAlternative) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SpeechRecognitionAlternative) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *SpeechRecognitionAlternative) GetWords() []*WordInfo {
	if x != nil {
		return x.Words
	}
	return nil
}

type WordInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Time offset relative to the beginning of the audio,
	// and corresponding to the start of the spoken word.
	// This field is only set if `enable_word_time_offsets=true` and only
	// in the top hypothesis.
	// This is an experimental feature and the accuracy of the time offset can
	// vary.
	StartAt int64 `protobuf:"varint,1,opt,name=start_at,json=startAt,proto3" json:"start_at,omitempty"`
	// Time offset relative to the beginning of the audio,
	// and corresponding to the end of the spoken word.
	// This field is only set if `enable_word_time_offsets=true` and only
	// in the top hypothesis.
	// This is an experimental feature and the accuracy of the time offset can
	// vary.
	Duration int64 `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	// The word corresponding to this set of information.
	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
	// The confidence estimate between 0.0 and 1.0. A higher number
	// indicates an estimated greater likelihood that the recognized words are
	// correct. This field is set only for the top alternative of a non-streaming
	// result or, of a streaming result where `is_final=true`.
	// This field is not guaranteed to be accurate and users should not rely on it
	// to be always provided.
	// The default of 0.0 is a sentinel value indicating `confidence` was not set.
	Confidence float32 `protobuf:"fixed32,4,opt,name=confidence,proto3" json:"confidence,omitempty"`
	// A distinct integer value is assigned for every speaker within
	// the audio. This field specifies which one of those speakers was detected to
	// have spoken this word. Value ranges from '1' to diarization_speaker_count.
	// speaker_tag is set if enable_speaker_diarization = 'true' and only in the
	// top alternative.
	SpeakerTag int32 `protobuf:"varint,5,opt,name=speaker_tag,json=speakerTag,proto3" json:"speaker_tag,omitempty"`
}

func (x *WordInfo) Reset() {
	*x = WordInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WordInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WordInfo) ProtoMessage() {}

func (x *WordInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WordInfo.ProtoReflect.Descriptor instead.
func (*WordInfo) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{4}
}

func (x *WordInfo) GetStartAt() int64 {
	if x != nil {
		return x.StartAt
	}
	return 0
}

func (x *WordInfo) GetDuration() int64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *WordInfo) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *WordInfo) GetConfidence() float32 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

func (x *WordInfo) GetSpeakerTag() int32 {
	if x != nil {
		return x.SpeakerTag
	}
	return 0
}

type DecoderConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SampleRate          int32                       `protobuf:"varint,1,opt,name=sample_rate,json=sampleRate,proto3" json:"sample_rate,omitempty"`
	Encoding            DecoderConfig_AudioEncoding `protobuf:"varint,2,opt,name=encoding,proto3,enum=online_decoder.DecoderConfig_AudioEncoding" json:"encoding,omitempty"`
	UseItn              *bool                       `protobuf:"varint,5,opt,name=use_itn,json=useItn,proto3,oneof" json:"use_itn,omitempty"`
	UseDisfluencyFilter *bool                       `protobuf:"varint,13,opt,name=use_disfluency_filter,json=useDisfluencyFilter,proto3,oneof" json:"use_disfluency_filter,omitempty"`
	UseProfanityFilter  *bool                       `protobuf:"varint,14,opt,name=use_profanity_filter,json=useProfanityFilter,proto3,oneof" json:"use_profanity_filter,omitempty"`
}

func (x *DecoderConfig) Reset() {
	*x = DecoderConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vito_stt_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DecoderConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecoderConfig) ProtoMessage() {}

func (x *DecoderConfig) ProtoReflect() protoreflect.Message {
	mi := &file_vito_stt_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecoderConfig.ProtoReflect.Descriptor instead.
func (*DecoderConfig) Descriptor() ([]byte, []int) {
	return file_vito_stt_client_proto_rawDescGZIP(), []int{5}
}

func (x *DecoderConfig) GetSampleRate() int32 {
	if x != nil {
		return x.SampleRate
	}
	return 0
}

func (x *DecoderConfig) GetEncoding() DecoderConfig_AudioEncoding {
	if x != nil {
		return x.Encoding
	}
	return DecoderConfig_ENCODING_UNSPECIFIED
}

func (x *DecoderConfig) GetUseItn() bool {
	if x != nil && x.UseItn != nil {
		return *x.UseItn
	}
	return false
}

func (x *DecoderConfig) GetUseDisfluencyFilter() bool {
	if x != nil && x.UseDisfluencyFilter != nil {
		return *x.UseDisfluencyFilter
	}
	return false
}

func (x *DecoderConfig) GetUseProfanityFilter() bool {
	if x != nil && x.UseProfanityFilter != nil {
		return *x.UseProfanityFilter
	}
	return false
}

var File_vito_stt_client_proto protoreflect.FileDescriptor

var file_vito_stt_client_proto_rawDesc = []byte{
	0x0a, 0x15, 0x76, 0x69, 0x74, 0x6f, 0x2d, 0x73, 0x74, 0x74, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x22, 0x98, 0x01, 0x0a, 0x0e, 0x44, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4a, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0d, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52,
	0x0c, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x13, 0x0a,
	0x11, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x98, 0x02, 0x0a, 0x0f, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x44, 0x0a, 0x07,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x12, 0x5b, 0x0a, 0x11, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0f,
	0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x4c, 0x0a, 0x0f, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x50, 0x45, 0x45, 0x43, 0x48, 0x5f, 0x45, 0x56, 0x45,
	0x4e, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x1b, 0x0a, 0x17, 0x45, 0x4e, 0x44, 0x5f, 0x4f, 0x46, 0x5f, 0x53, 0x49, 0x4e, 0x47, 0x4c,
	0x45, 0x5f, 0x55, 0x54, 0x54, 0x45, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x22, 0xde, 0x01,
	0x0a, 0x1a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x63, 0x6f, 0x67,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x50, 0x0a, 0x0c,
	0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x2e, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x52, 0x0c, 0x61, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x73, 0x12, 0x19,
	0x0a, 0x08, 0x69, 0x73, 0x5f, 0x66, 0x69, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x69, 0x73, 0x46, 0x69, 0x6e, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x73, 0x74,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x22, 0x82,
	0x01, 0x0a, 0x1c, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x65, 0x63, 0x6f, 0x67, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65,
	0x6e, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f,
	0x64, 0x65, 0x72, 0x2e, 0x57, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x77, 0x6f,
	0x72, 0x64, 0x73, 0x22, 0x96, 0x01, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x73, 0x74, 0x61, 0x72, 0x74, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72, 0x54, 0x61, 0x67, 0x22, 0xd5, 0x03, 0x0a,
	0x0d, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x47, 0x0a, 0x08, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2b, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x08,
	0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x5f,
	0x69, 0x74, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x49, 0x74, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x37, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x69,
	0x73, 0x66, 0x6c, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x13, 0x75, 0x73, 0x65, 0x44, 0x69, 0x73, 0x66,
	0x6c, 0x75, 0x65, 0x6e, 0x63, 0x79, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12,
	0x35, 0x0a, 0x14, 0x75, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79,
	0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52,
	0x12, 0x75, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x46, 0x69, 0x6c,
	0x74, 0x65, 0x72, 0x88, 0x01, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x0d, 0x41, 0x75, 0x64, 0x69, 0x6f,
	0x45, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e, 0x43, 0x4f,
	0x44, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x49, 0x4e, 0x45, 0x41, 0x52, 0x31, 0x36, 0x10, 0x01,
	0x12, 0x07, 0x0a, 0x03, 0x57, 0x41, 0x56, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x4c, 0x41,
	0x43, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x55, 0x4c, 0x41, 0x57, 0x10, 0x04, 0x12, 0x08,
	0x0a, 0x04, 0x41, 0x4c, 0x41, 0x57, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4d, 0x52, 0x10,
	0x06, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x4d, 0x52, 0x5f, 0x57, 0x42, 0x10, 0x07, 0x12, 0x0c, 0x0a,
	0x08, 0x4f, 0x47, 0x47, 0x5f, 0x4f, 0x50, 0x55, 0x53, 0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x4f,
	0x50, 0x55, 0x53, 0x10, 0x09, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x69, 0x74,
	0x6e, 0x42, 0x18, 0x0a, 0x16, 0x5f, 0x75, 0x73, 0x65, 0x5f, 0x64, 0x69, 0x73, 0x66, 0x6c, 0x75,
	0x65, 0x6e, 0x63, 0x79, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x17, 0x0a, 0x15, 0x5f,
	0x75, 0x73, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x61, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x32, 0x60, 0x0a, 0x0d, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x44, 0x65,
	0x63, 0x6f, 0x64, 0x65, 0x72, 0x12, 0x4f, 0x0a, 0x06, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x1e, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x63, 0x6f, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x28, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x69, 0x74, 0x6f, 0x2d, 0x61, 0x69, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x74, 0x3b, 0x73, 0x74, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vito_stt_client_proto_rawDescOnce sync.Once
	file_vito_stt_client_proto_rawDescData = file_vito_stt_client_proto_rawDesc
)

func file_vito_stt_client_proto_rawDescGZIP() []byte {
	file_vito_stt_client_proto_rawDescOnce.Do(func() {
		file_vito_stt_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_vito_stt_client_proto_rawDescData)
	})
	return file_vito_stt_client_proto_rawDescData
}

var file_vito_stt_client_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_vito_stt_client_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_vito_stt_client_proto_goTypes = []interface{}{
	(DecoderResponse_SpeechEventType)(0), // 0: online_decoder.DecoderResponse.SpeechEventType
	(DecoderConfig_AudioEncoding)(0),     // 1: online_decoder.DecoderConfig.AudioEncoding
	(*DecoderRequest)(nil),               // 2: online_decoder.DecoderRequest
	(*DecoderResponse)(nil),              // 3: online_decoder.DecoderResponse
	(*StreamingRecognitionResult)(nil),   // 4: online_decoder.StreamingRecognitionResult
	(*SpeechRecognitionAlternative)(nil), // 5: online_decoder.SpeechRecognitionAlternative
	(*WordInfo)(nil),                     // 6: online_decoder.WordInfo
	(*DecoderConfig)(nil),                // 7: online_decoder.DecoderConfig
}
var file_vito_stt_client_proto_depIdxs = []int32{
	7, // 0: online_decoder.DecoderRequest.streaming_config:type_name -> online_decoder.DecoderConfig
	4, // 1: online_decoder.DecoderResponse.results:type_name -> online_decoder.StreamingRecognitionResult
	0, // 2: online_decoder.DecoderResponse.speech_event_type:type_name -> online_decoder.DecoderResponse.SpeechEventType
	5, // 3: online_decoder.StreamingRecognitionResult.alternatives:type_name -> online_decoder.SpeechRecognitionAlternative
	6, // 4: online_decoder.SpeechRecognitionAlternative.words:type_name -> online_decoder.WordInfo
	1, // 5: online_decoder.DecoderConfig.encoding:type_name -> online_decoder.DecoderConfig.AudioEncoding
	2, // 6: online_decoder.OnlineDecoder.Decode:input_type -> online_decoder.DecoderRequest
	3, // 7: online_decoder.OnlineDecoder.Decode:output_type -> online_decoder.DecoderResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_vito_stt_client_proto_init() }
func file_vito_stt_client_proto_init() {
	if File_vito_stt_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vito_stt_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecoderRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vito_stt_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecoderResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vito_stt_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamingRecognitionResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vito_stt_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeechRecognitionAlternative); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vito_stt_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WordInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vito_stt_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DecoderConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vito_stt_client_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DecoderRequest_StreamingConfig)(nil),
		(*DecoderRequest_AudioContent)(nil),
	}
	file_vito_stt_client_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vito_stt_client_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vito_stt_client_proto_goTypes,
		DependencyIndexes: file_vito_stt_client_proto_depIdxs,
		EnumInfos:         file_vito_stt_client_proto_enumTypes,
		MessageInfos:      file_vito_stt_client_proto_msgTypes,
	}.Build()
	File_vito_stt_client_proto = out.File
	file_vito_stt_client_proto_rawDesc = nil
	file_vito_stt_client_proto_goTypes = nil
	file_vito_stt_client_proto_depIdxs = nil
}
