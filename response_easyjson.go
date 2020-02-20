// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package fcm

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson6ff3ac1dDecodeGithubComHumansNetFcm(in *jlexer.Lexer, out *Result) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "message_id":
			out.MessageID = string(in.String())
		case "registration_id":
			out.RegistrationID = string(in.String())
		case "error":
			out.Error = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6ff3ac1dEncodeGithubComHumansNetFcm(out *jwriter.Writer, in Result) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"message_id\":"
		out.RawString(prefix[1:])
		out.String(string(in.MessageID))
	}
	{
		const prefix string = ",\"registration_id\":"
		out.RawString(prefix)
		out.String(string(in.RegistrationID))
	}
	{
		const prefix string = ",\"error\":"
		out.RawString(prefix)
		out.String(string(in.Error))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Result) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComHumansNetFcm(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Result) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComHumansNetFcm(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Result) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComHumansNetFcm(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Result) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComHumansNetFcm(l, v)
}
func easyjson6ff3ac1dDecodeGithubComHumansNetFcm1(in *jlexer.Lexer, out *Response) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "multicast_id":
			out.MulticastID = int64(in.Int64())
		case "success":
			out.Success = int(in.Int())
		case "failure":
			out.Failure = int(in.Int())
		case "canonical_ids":
			out.CanonicalIDs = int(in.Int())
		case "error_code":
			out.StatusCode = int(in.Int())
		case "results":
			if in.IsNull() {
				in.Skip()
				out.Results = nil
			} else {
				in.Delim('[')
				if out.Results == nil {
					if !in.IsDelim(']') {
						out.Results = make([]Result, 0, 1)
					} else {
						out.Results = []Result{}
					}
				} else {
					out.Results = (out.Results)[:0]
				}
				for !in.IsDelim(']') {
					var v1 Result
					(v1).UnmarshalEasyJSON(in)
					out.Results = append(out.Results, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson6ff3ac1dEncodeGithubComHumansNetFcm1(out *jwriter.Writer, in Response) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"multicast_id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.MulticastID))
	}
	{
		const prefix string = ",\"success\":"
		out.RawString(prefix)
		out.Int(int(in.Success))
	}
	{
		const prefix string = ",\"failure\":"
		out.RawString(prefix)
		out.Int(int(in.Failure))
	}
	{
		const prefix string = ",\"canonical_ids\":"
		out.RawString(prefix)
		out.Int(int(in.CanonicalIDs))
	}
	{
		const prefix string = ",\"error_code\":"
		out.RawString(prefix)
		out.Int(int(in.StatusCode))
	}
	{
		const prefix string = ",\"results\":"
		out.RawString(prefix)
		if in.Results == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Results {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Response) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson6ff3ac1dEncodeGithubComHumansNetFcm1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Response) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson6ff3ac1dEncodeGithubComHumansNetFcm1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Response) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson6ff3ac1dDecodeGithubComHumansNetFcm1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Response) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson6ff3ac1dDecodeGithubComHumansNetFcm1(l, v)
}
