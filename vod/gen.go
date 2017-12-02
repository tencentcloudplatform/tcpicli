// +build ignore

package main

import (
	. "."
	"github.com/tencentcloudplatform/tcpicli/autogen"
)

type Pkg struct{}

func (p *Pkg) DoAction(action string, query ...string) ([]byte, error) {
	return DoAction(action, query...)
}

func main() {
	region := "Region=bj"                                                             // common input param
	className := "className=apitest"                                                  // CreateClass etc
	classNameMod := "className=apitestmod"                                            // ModifyClass
	fileName := "fileName=apitest"                                                    // DescribeVodPlayInfo // need to statically bootstrap because cannot upload files until COS API is ironed out
	watermarkName := "name=apitest"                                                   // CreateWatermarkTemplate
	watermarkUrl := "url=tb1.bdstatic.com/tb/cms/ngmis/images/file_1512126542603.jpg" // CreateWatermarkTemplate
	width := "width=100px"                                                            // CreateWatermarkTemplate
	watermarkType := "type=image"                                                     // CreatewatermarkTemplate
	widthMod := "width=110px"                                                         // UpdateWatermarkTemplate
	transcodeName := "name=apitest"                                                   // QueryTranscodeTemplateList
	transcodeContainer := "container=mp4"                                             // CreateTranscodeTemplate
	transcodeComment := "comment=apitest"                                             // CreateTranscodeTemplate
	transcodeVideoCodec := "video.codec=libx264"                                      // CreateTranscodeTemplate
	transcodeFps := "video.fps=45"                                                    // CreateTranscodeTemplate
	transcodeWidth := "video.width=1280"                                              // CreateTranscodeTemplate
	transcodeHeight := "video.height=720"                                             // CreateTranscodeTemplate
	transcodeVideoBitrate := "video.bitrate=512"                                      // CreateTranscodeTemplate
	transcodeAudioCodec := "audio.codec=libfdk_aac"                                   // CreateTranscodeTemplate
	transcodeAudioBitrate := "audio.bitrate=256"                                      // CreateTranscodeTemplate
	transcodeSoundSystem := "audio.soundSystem=2"                                     // CreateTranscodeTemplate
	transcodeAudioSampleRate := "audio.sampleRate=44100"                              // CreateTranscodeTemplate
	transcodeCommentMod := "comment=apitestmod"                                       // UpdateTranscodeTemplate
	tags := "tags.1=apitest"                                                          // CreateVodTags // 1 index -_-
	vodDescriptionMod := "fileIntro=apitestmod"                                       // ModifyVodInfo // fileIntro... *sigh*

	gen := &autogen.Gen{
		DocRoot: "https://cloud.tencent.com/document/api/",
		Seq: []string{
			// -- hold on upload stuff because it has COS API dependency --
			// "ApplyUpload",
			// "CommitUpload",
			// "MultiPullVodFile",
			// "ApplyUploadWatermark",

			// "RunProcedure",
			// "ProcessFile",
			// "ConvertVodFile",
			// "ClipVideo",
			// "ConcatVideo",
			// "CreateSnapshotByTimeOffset",
			// "CreateImageSprite",
			"DescribeVodPlayInfo",
			`SET fileId=tcpicli -f "{{range .FileSet}}{{.FileID}}{{end}}" vod DescribeVodPlayInfo ` + region + " " + fileName,
			`SET vid=echo $fileId`, // stupid -- two variables, one value ;)
			"GetVideoInfo",
			"DescribeRecordPlayInfo",
			"CreateVodTags",
			"DeleteVodTags",
			// "DeleteVodFile",
			// "ModifyVodInfo",
			"CreateClass",
			"DescribeAllClass",
			"DescribeClass",
			// inspect for vod sucks because API doesn't allow you to filter by className or name or anything descriptive; only barfs all classes at you. there is an immortal default class, therefore, to filter I need to use bash...
			`SET classId=tcpicli vod DescribeClass ` + region + " | grep -B1 $(echo " + className + " | cut -d '=' -f2)" + ` | grep -v name | awk '{ print $2 }' | tr -d "\","`,
			"ModifyClass",
			"DeleteClass",
			// "PullEvent",
			// "ConfirmEvent",
			// "GetTaskList",
			// "GetTaskInfo",
			// "RedoTask",
			"CreateTranscodeTemplate",
			"QueryTranscodeTemplateList",
			`SET transcodeDefinition=tcpicli -f "{{range .Data}}{{.Definition}}{{end}}" vod QueryTranscodeTemplateList`,
			"QueryTranscodeTemplate",
			"UpdateTranscodeTemplate",
			"DeleteTranscodeTemplate",
			"CreateWatermarkTemplate", // can create from http :D
			"QueryWatermarkTemplateList",
			`SET watermarkDefinition=tcpicli -f "{{.Zero.Definition}}" vod QueryWatermarkTemplateList ` + watermarkName,
			"QueryWatermarkTemplate",
			"UpdateWatermarkTemplate",
			"DeleteWatermarkTemplate",
			"DescribeDrmDataKey",
			// "DescribeVodStorage",
			// "GetPlayStatLogList",
		},
		FuncMap: map[string][]string{
			// -- hold on upload stuff because it has COS API dependency --
			// "ApplyUpload":                []string{"266/9756"},
			// "CommitUpload":               []string{"266/9757"},
			// "MultiPullVodFile":           []string{"266/7817"},
			// "ApplyUploadWatermark":       []string{"266/11607"},

			"RunProcedure":               []string{"266/11030"},
			"ProcessFile":                []string{"266/9642"},
			"ConvertVodFile":             []string{"266/7822"},
			"ClipVideo":                  []string{"266/10156"},
			"ConcatVideo":                []string{"266/7821"},
			"CreateSnapshotByTimeOffset": []string{"266/8102"},
			"CreateImageSprite":          []string{"266/8101"},
			"GetVideoInfo": []string{"266/8586",
				region,
				"fileId=$fileId",
			},
			"DescribeVodPlayInfo": []string{"266/7825",
				region,
				fileName,
			},
			"DescribeRecordPlayInfo": []string{"266/8227",
				region,
				"vid=$vid",
			},
			"CreateVodTags": []string{"266/7826",
				region,
				tags,
				"fileId=$fileId",
			},
			"DeleteVodTags": []string{"266/7827",
				region,
				tags,
				"fileId=$fileId",
			},
			"DeleteVodFile": []string{"266/7838"},
			"ModifyVodInfo": []string{"266/7828",
				region,
				vodDescriptionMod,
				"fileId=$fileId",
			},
			"CreateClass": []string{"266/7812",
				region,
				className,
			},
			"DescribeAllClass": []string{"266/7813",
				region,
			},
			"DescribeClass": []string{"266/7814",
				region,
			},
			"ModifyClass": []string{"266/7815",
				region,
				classNameMod,
				"classId=$classId",
			},
			"DeleteClass": []string{"266/7816",
				region,
				"classId=$classId",
			},
			"PullEvent":    []string{"266/7818"},
			"ConfirmEvent": []string{"266/7819"},
			"GetTaskList":  []string{"266/11722"},
			"GetTaskInfo":  []string{"266/11724"},
			"RedoTask":     []string{"266/11725"},
			"CreateTranscodeTemplate": []string{"266/9910",
				region,
				transcodeName,
				transcodeContainer,
				transcodeComment,
				transcodeVideoCodec,
				transcodeFps,
				transcodeWidth,
				transcodeHeight,
				transcodeVideoBitrate,
				transcodeAudioCodec,
				transcodeAudioBitrate,
				transcodeSoundSystem,
				transcodeAudioSampleRate,
			},
			"QueryTranscodeTemplateList": []string{"266/9913",
				region,
				transcodeName,
			},
			"QueryTranscodeTemplate": []string{"266/9912",
				region,
				"definition=$transcodeDefinition",
			},
			"UpdateTranscodeTemplate": []string{"266/9911",
				region,
				transcodeCommentMod,
				"definition=$transcodeDefinition",
			},
			"DeleteTranscodeTemplate": []string{"266/9914",
				region,
				"definition=$transcodeDefinition",
			},
			"CreateWatermarkTemplate": []string{"266/11599",
				region,
				watermarkName,
				watermarkUrl,
				width,
				watermarkType,
			},
			"QueryWatermarkTemplateList": []string{"266/11608",
				region,
				watermarkName,
			},
			"QueryWatermarkTemplate": []string{"266/11606",
				region,
				"definition=$watermarkDefinition",
			},
			"UpdateWatermarkTemplate": []string{"266/11605",
				region,
				widthMod,
				"definition=$watermarkDefinition",
			},
			"DeleteWatermarkTemplate": []string{"266/11604",
				region,
				"definition=$watermarkDefinition",
			},
			"DescribeDrmDataKey": []string{"266/9643"},
			"DescribeVodStorage": []string{"266/10012"},
			"GetPlayStatLogList": []string{"266/12624"},
		},
		PkgName: "vod",
		Pkg:     new(Pkg),
	}

	gen.Run()
}
