package conv

import (
	"fmt"
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"github.com/zhangyiming748/ConvertVideo/replace"
	"github.com/zhangyiming748/ConvertVideo/util"
	"log/slog"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	//AudioBook = "1.54" //等效audition的65%
	AudioBook = "1.43" //等效audition的70%

)

/*
加速单个音频的完整函数
*/
func SpeedupAudio(in mediainfo.BasicInfo, speed string) {
	dst := strings.Join([]string{in.PurgePath, "speed"}, string(os.PathSeparator)) //目标文件目录
	os.Mkdir(dst, 0777)
	fname := replace.ForFileName(in.PurgeName)
	fname = strings.Join([]string{fname, "mp3"}, ".")
	slog.Debug("补全后的 fname", slog.String("fname", fname))
	out := strings.Join([]string{dst, fname}, string(os.PathSeparator))
	slog.Debug("io", slog.String("输入文件", in.FullPath), slog.String("输出文件", out))
	//跳过已经加速的文件夹
	if strings.Contains(in.FullPath, "speed") {
		return
	}
	speedUp(in.FullPath, out, speed)
}

/*
仅使用输入输出和加速参数执行命令
*/
func speedUp(in, out string, speed string) {
	ff := audition2ffmpeg(speed)
	atempo := strings.Join([]string{"atempo", ff}, "=")
	filter := strings.Join([]string{atempo, "volume=3.0"}, ",")
	cmd := exec.Command("ffmpeg", "-i", in, "-filter:a", filter, "-vn", "-ac", "1", "-map_metadata", "-1", out)
	util.ExecCommand(cmd, in)

}

/*
获取一个等效adobe audition 的 混缩
*/
func audition2ffmpeg(speed string) string {
	audition, err := strconv.ParseFloat(speed, 64)
	if err != nil {
		slog.Warn("解析加速参数错误,退出程序", slog.String("错误原文", fmt.Sprint(err)))
		os.Exit(1)
	}
	param := 100 / audition
	slog.Debug("转换后的原始参数", slog.Float64("param", param))
	final := fmt.Sprintf("%.2f", param)
	slog.Debug("保留两位小数的原始参数", slog.String("final", final))
	return final
}
