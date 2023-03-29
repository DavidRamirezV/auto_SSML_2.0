package functions

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	//go get
	//texttospeechpb "google.golang.org/genproto/googleapis/cloud/texttospeech/v1"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	texttospeechpb "cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

func Synthethize(ssml string, filename string) {
	// Instantiates a client.
	ctx := context.Background()

	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	req := texttospeechpb.SynthesizeSpeechRequest{
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Ssml{Ssml: ssml},
		},
		// Note: the voice can also be specified by name.
		// Names of voices can be retrieved with client.ListVoices().
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "es-US",
			Name:         "es-US-Neural2-B",
			SsmlGender:   texttospeechpb.SsmlVoiceGender_MALE,
		},
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	//filename
	filename_final := ""
	if filename == "" {
		filename_final = "synthethized.mp3"
	} else {
		filename_final = strings.Replace(filename, ".", "_", -1) + ".mp3"
	}

	// The resp's AudioContent is binary.

	err = ioutil.WriteFile("files/"+filename_final, resp.AudioContent, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nArchivo '%v' guardado en la carpeta /src/files/\n\n", filename_final)

}
