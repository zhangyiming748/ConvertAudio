package conv

import (
	"github.com/zhangyiming748/ConvertVideo/mediainfo"
	"log/slog"
	"os"
	"os/exec"
	"strings"
)

func SingleAAC(in mediainfo.BasicInfo) {
	out := strings.Replace(in.FullPath, in.PurgeExt, "aac", 1)
	cmd := exec.Command("ffmpeg", "-i", in.FullPath, "-ac", "1", out)
	output, err := cmd.CombinedOutput()
	if err != nil {
		slog.Error("命令执行失败", slog.String("命令原文", cmd.String()), slog.String("命令输出", string(output)), slog.String("错误原文", err.Error()))
		return
	} else {
		slog.Info("命令成功执行", slog.String("输出", string(output)))
		os.Remove(in.FullPath)
	}
}
