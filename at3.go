package main

import (
	"bufio"
	"github.com/Songmu/axslogparser"
	"strconv"
	"io"
	"encoding/csv"
)

type At3 struct {
	parser *axslogparser.Apache
}

func NewAt3() *At3 {
	return &At3{
		&axslogparser.Apache{},
	}
}

func (at3 *At3) Transform(in io.Reader, out io.Writer) error {
	writer := csv.NewWriter(out)

	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		l, err := at3.parser.Parse(line)
		if err != nil {
			return err
		}

		record := parseLog(l)
		writer.Write(record)
	}
	writer.Flush()

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func parseLog(log *axslogparser.Log) []string {
	status := strconv.Itoa(log.Status)
	result := []string{
		log.Host,
		log.RemoteUser,
		log.User,
		log.Time.String(),
		log.Method,
		log.RequestURI,
		log.Protocol,
		status,
		log.Referer,
		log.UserAgent,
	}
	return result
}
