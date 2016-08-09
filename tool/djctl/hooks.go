package main

import (
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/gravitational/rigging"
	"github.com/gravitational/trace"
)

func createDB(name string) error {
	jobName := "stolon-createdb"
	template := `
apiVersion: batch/v1
kind: Job
metadata:
  name: ` + jobName + `
spec:
  template:
    metadata:
      name: ` + jobName + `
    spec:
      containers:
      - name: createdb
        imagePullPolicy: Always
        image: apiserver:5000/stolon:latest
        args: ["db", "create", "` + name + `"]
        env:
          - name: CTL
            value: "true"
          - name: STOLONCTL_DB_USERNAME
            value: "stolon"
          - name: STOLONCTL_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: stolon
                key: password
      restartPolicy: OnFailure
`

	log.Infof("Create job for creating database %s", name)
	out, err := rigging.FromStdIn(rigging.ActionCreate, template)
	if err != nil {
		log.Errorf("%s", string(out))
		return trace.Wrap(err)
	}

	log.Infof("waiting for job %s", jobName)
	err = rigging.WaitForJobSuccess(jobName, 1*time.Minute)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

func backupDB(name string) error {
	jobName := "stolon-backup"
	template := `
apiVersion: batch/v1
kind: Job
metadata:
  name: ` + jobName + `
spec:
  template:
    metadata:
      name: ` + jobName + `
    spec:
      containers:
      - name: backup
        imagePullPolicy: Always
        image: apiserver:5000/stolon:latest
        args: ["db", "backup", "` + name + `", "/backups"]
        env:
          - name: CTL
            value: "true"
          - name: STOLONCTL_DB_USERNAME
            value: "stolon"
          - name: STOLONCTL_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: stolon
                key: password
        volumeMounts:
          - mountPath: /backups
            name: backups
      volumes:
        - name: backups
          hostPath:
            path: /var/backups
      restartPolicy: OnFailure
`

	log.Infof("Create job for backup database %s", name)
	out, err := rigging.FromStdIn(rigging.ActionCreate, template)
	if err != nil {
		log.Errorf("%s", string(out))
		return trace.Wrap(err)
	}

	log.Infof("waiting for job %s", jobName)
	err = rigging.WaitForJobSuccess(jobName, 1*time.Minute)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}

func restoreDB(name, src string) {
	jobName := "stolon-restore"
	template := `
apiVersion: batch/v1
kind: Job
metadata:
  name: ` + jobName + `
spec:
  template:
    metadata:
      name: ` + jobName + `
    spec:
      containers:
      - name: restore
        imagePullPolicy: Always
        image: apiserver:5000/stolon:latest
        args: ["db", "restore", "` + src + `"]
        env:
          - name: CTL
            value: "true"
          - name: STOLONCTL_DB_USERNAME
            value: "stolon"
          - name: STOLONCTL_DB_PASSWORD
            valueFrom:
              secretKeyRef:
                name: stolon
                key: password
        volumeMounts:
          - mountPath: /backups
            name: backups
      volumes:
        - name: backups
          hostPath:
            path: /var/backups
      restartPolicy: OnFailure
`

	log.Infof("Create job for backup database %s", name)
	out, err := rigging.FromStdIn(rigging.ActionCreate, template)
	if err != nil {
		log.Errorf("%s", string(out))
		return trace.Wrap(err)
	}

	log.Infof("waiting for job %s", jobName)
	err = rigging.WaitForJobSuccess(jobName, 1*time.Minute)
	if err != nil {
		return trace.Wrap(err)
	}

	return nil
}
