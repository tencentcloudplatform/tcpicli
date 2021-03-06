package main

import (
	"fmt"
	"github.com/tencentcloudplatform/tcpicli/vod"
	"github.com/urfave/cli"
)

var (
	funcVod []cli.Command = []cli.Command{
		{
			Name:   "do",
			Usage:  "do action and output raw json response.",
			Action: VodDoAction,
		},
		{
			Name:        "CreateClass",
			Usage:       "Create VOD classification.",
			Action:      VodCreateClass,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7812",
		},
		{
			Name:        "DescribeClass",
			Usage:       "Describe VOD classification.",
			Action:      VodDescribeClass,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7814",
		},
		{
			Name:        "DescribeAllClass",
			Usage:       "Describe all VOD classification.",
			Action:      VodDescribeAllClass,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7813",
		},
		{
			Name:        "ModifyClass",
			Usage:       "Modify VOD classification.",
			Action:      VodModifyClass,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7815",
		},
		{
			Name:        "DeleteClass",
			Usage:       "Delete VOD classification.",
			Action:      VodDeleteClass,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7816",
		},
		{
			Name:        "DescribeVodPlayInfo",
			Usage:       "Returns videofile information.",
			Action:      VodDescribeVodPlayInfo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7825",
		},
		{
			Name:        "DescribeRecordPlayInfo",
			Usage:       "Returns information.",
			Action:      VodDescribeRecordPlayInfo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/8227",
		},
		{
			Name:        "GetVideoInfo",
			Usage:       "Returns videofile information.",
			Action:      VodGetVideoInfo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/8586",
		},
		{
			Name:        "QueryWatermarkTemplateList",
			Usage:       "Returns list of watermark templates.",
			Action:      VodQueryWatermarkTemplateList,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11608",
		},
		{
			Name:        "QueryWatermarkTemplate",
			Usage:       "Returns information on watermark template",
			Action:      VodQueryWatermarkTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11606",
		},
		{
			Name:        "CreateWatermarkTemplate",
			Usage:       "Creates a new watermark template",
			Action:      VodCreateWatermarkTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11599",
		},
		{
			Name:        "UpdateWatermarkTemplate",
			Usage:       "Updates a watermark template",
			Action:      VodUpdateWatermarkTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11605",
		},
		{
			Name:        "DeleteWatermarkTemplate",
			Usage:       "Deletes a watermark template",
			Action:      VodDeleteWatermarkTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11604",
		},
		{
			Name:        "QueryTranscodeTemplateList",
			Usage:       "Queries list of transcode templates",
			Action:      VodQueryTranscodeTemplateList,
			Description: "referrer: https://cloud.tencent.com/document/api/266/9913",
		},
		{
			Name:        "CreateTranscodeTemplate",
			Usage:       "Creates a new template",
			Action:      VodCreateTranscodeTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/9910",
		},
		{
			Name:        "QueryTranscodeTemplate",
			Usage:       "Querys a particular transcode template",
			Action:      VodQueryTranscodeTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/9912",
		},
		{
			Name:        "DeleteTranscodeTemplate",
			Usage:       "Deletes a transcode template",
			Action:      VodDeleteTranscodeTemplate,
			Description: "referrer: https://cloud.tencent.com/document/api/266/9914",
		},
		{
			Name:        "CreateVodTags",
			Usage:       "Creates a tag on a videofile",
			Action:      VodCreateVodTags,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7826",
		},
		{
			Name:        "DeleteVodTags",
			Usage:       "Deletes a tag on a videofile",
			Action:      VodDeleteVodTags,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7827",
		},
		{
			Name:        "ModifyVodInfo",
			Usage:       "Deletes a tag on a videofile",
			Action:      VodModifyVodInfo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7828",
		},
		{
			Name:        "DescribeVodStorage",
			Usage:       "Describes VOD storage",
			Action:      VodDescribeVodStorage,
			Description: "referrer: https://cloud.tencent.com/document/api/266/10012",
		},
		{
			Name:        "GetPlayStatLogList",
			Usage:       "Returns log data",
			Action:      VodGetPlayStatLogList,
			Description: "referrer: https://cloud.tencent.com/document/api/266/12624",
		},
		{
			Name:        "ProcessFile",
			Usage:       "Processes VOD file.",
			Action:      VodProcessFile,
			Description: "referrer: https://cloud.tencent.com/document/api/266/9642",
		},
		{
			Name:        "GetTaskList",
			Usage:       "Returns VOD Task List.",
			Action:      VodGetTaskList,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11722",
		},
		{
			Name:        "ConvertVodFile",
			Usage:       "Converts VOD file",
			Action:      VodConvertVodFile,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7822",
		},
		{
			Name:        "ClipVideo",
			Usage:       "Creates cut of video and writes to a new file",
			Action:      VodClipVideo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/10156",
		},
		{
			Name:        "GetTaskInfo",
			Usage:       "Specifies info about specific task",
			Action:      VodGetTaskInfo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11724",
		},
		{
			Name:        "RedoTask",
			Usage:       "Redoes task by vodTaskId",
			Action:      VodRedoTask,
			Description: "referrer: https://cloud.tencent.com/document/api/266/11725",
		},
		{
			Name:        "ConcatVideo",
			Usage:       "Splices uploaded videos together",
			Action:      VodConcatVideo,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7821",
		},
		{
			Name:        "DeleteVodFile",
			Usage:       "Deletes specified VOD file.",
			Action:      VodDeleteVodFile,
			Description: "referrer: https://cloud.tencent.com/document/api/266/7838",
		},
	}
)

func VodDoAction(c *cli.Context) error {
	resp, err := vod.DoAction(c.Args().First(), c.Args().Tail()...)
	if err != nil {
		return err
	}
	fmt.Println(string(resp))
	return nil
}
func VodCreateClass(c *cli.Context) error {
	resp, err := vod.CreateClass(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDescribeClass(c *cli.Context) error {
	resp, err := vod.DescribeClass(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDescribeAllClass(c *cli.Context) error {
	resp, err := vod.DescribeAllClass(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodModifyClass(c *cli.Context) error {
	resp, err := vod.ModifyClass(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDeleteClass(c *cli.Context) error {
	resp, err := vod.DeleteClass(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDescribeVodPlayInfo(c *cli.Context) error {
	resp, err := vod.DescribeVodPlayInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDescribeRecordPlayInfo(c *cli.Context) error {
	resp, err := vod.DescribeRecordPlayInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodGetVideoInfo(c *cli.Context) error {
	resp, err := vod.GetVideoInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodQueryWatermarkTemplateList(c *cli.Context) error {
	resp, err := vod.QueryWatermarkTemplateList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodQueryWatermarkTemplate(c *cli.Context) error {
	resp, err := vod.QueryWatermarkTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodCreateWatermarkTemplate(c *cli.Context) error {
	resp, err := vod.CreateWatermarkTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodUpdateWatermarkTemplate(c *cli.Context) error {
	resp, err := vod.UpdateWatermarkTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDeleteWatermarkTemplate(c *cli.Context) error {
	resp, err := vod.DeleteWatermarkTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodQueryTranscodeTemplateList(c *cli.Context) error {
	resp, err := vod.QueryTranscodeTemplateList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodCreateTranscodeTemplate(c *cli.Context) error {
	resp, err := vod.CreateTranscodeTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodQueryTranscodeTemplate(c *cli.Context) error {
	resp, err := vod.QueryTranscodeTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDeleteTranscodeTemplate(c *cli.Context) error {
	resp, err := vod.DeleteTranscodeTemplate(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodCreateVodTags(c *cli.Context) error {
	resp, err := vod.CreateVodTags(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDeleteVodTags(c *cli.Context) error {
	resp, err := vod.DeleteVodTags(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodModifyVodInfo(c *cli.Context) error {
	resp, err := vod.ModifyVodInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDescribeVodStorage(c *cli.Context) error {
	resp, err := vod.DescribeVodStorage(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodGetPlayStatLogList(c *cli.Context) error {
	resp, err := vod.GetPlayStatLogList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodProcessFile(c *cli.Context) error {
	resp, err := vod.ProcessFile(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodGetTaskList(c *cli.Context) error {
	resp, err := vod.GetTaskList(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodConvertVodFile(c *cli.Context) error {
	resp, err := vod.ConvertVodFile(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodClipVideo(c *cli.Context) error {
	resp, err := vod.ClipVideo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodGetTaskInfo(c *cli.Context) error {
	resp, err := vod.GetTaskInfo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodRedoTask(c *cli.Context) error {
	resp, err := vod.RedoTask(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodConcatVideo(c *cli.Context) error {
	resp, err := vod.ConcatVideo(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
func VodDeleteVodFile(c *cli.Context) error {
	resp, err := vod.DeleteVodFile(c.Args()...)
	if err != nil {
		return err
	}
	r, err := resp.String(formatOut)
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}
