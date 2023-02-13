package main

import (
	"os"
	"os/exec"
	"strings"

	"github.com/docker/go-plugins-helpers/volume"
	mountedvolume "github.com/marcelo-ochoa/docker-volume-plugins/mounted-volume"
)

type s3fsDriver struct {
	defaultS3fsopts string
	mountedvolume.Driver
}

func (p *s3fsDriver) Validate(req *volume.CreateRequest) error {

	return nil
}

func (p *s3fsDriver) MountOptions(req *volume.CreateRequest) []string {

	s3fsopts, s3fsoptsInOpts := req.Options["s3fsopts"]

	var s3fsoptsArray []string
	if s3fsoptsInOpts {
		s3fsoptsArray = append(s3fsoptsArray, strings.Split(s3fsopts, ",")...)
	} else {
		s3fsoptsArray = append(s3fsoptsArray, strings.Split(p.defaultS3fsopts, ",")...)
	}
	s3fsoptsArray = AppendBucketOptionsByVolumeName(s3fsoptsArray, req.Name)

	return []string{"-o", strings.Join(s3fsoptsArray, ",")}
}

func (p *s3fsDriver) PreMount(req *volume.MountRequest) error {
	return nil
}

func (p *s3fsDriver) PostMount(req *volume.MountRequest) {
}

// AppendBucketOptionsByVolumeName appends the command line arguments into the current argument list given the volume name
func AppendBucketOptionsByVolumeName(args []string, volumeName string) []string {
	parts := strings.SplitN(volumeName, "/", 2)
	if len(parts) == 2 {
		return append(args, "bucket="+parts[0]+":/"+parts[1])
	}
	return append(args, "bucket="+parts[0])
}

func buildDriver() *s3fsDriver {
	defaultsopts := os.Getenv("DEFAULT_S3FSOPTS")
	d := &s3fsDriver{
		Driver:          *mountedvolume.NewDriver("s3fs", false, "s3fs", "local"),
		defaultS3fsopts: defaultsopts,
	}
	d.Init(d)
	return d
}

func spawnSyslog() {
	cmd := exec.Command("rsyslogd", "-n")
	cmd.Start()
}

func main() {
	spawnSyslog()
	//log.SetFlags(0)
	d := buildDriver()
	defer d.Close()
	d.ServeUnix()
}
