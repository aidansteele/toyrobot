package toyrobot

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var placeRegex = regexp.MustCompile(`(?i)^\s*place\s+(\d+)\s*,\s*(\d+)\s*,\s*(north|south|east|west)\s*$`)
var moveRegex = regexp.MustCompile(`(?i)^\s*move\s*$`)
var leftRegex = regexp.MustCompile(`(?i)^\s*left\s*$`)
var rightRegex = regexp.MustCompile(`(?i)^\s*right\s*$`)
var reportRegex = regexp.MustCompile(`(?i)^\s*report\s*$`)

// CommandReceiver is the interface that a type must conform to if it wishes
// to receive parsed input commands from the Parser class. Each method
// corresponds to one of the invocations specified in the Toy Robot Spec, with
// the except of Invalid(), which is invoked upon encountering a line that
// doesn't match any recognised command.
type CommandReceiver interface {
	Place(p position)
	Move()
	Left()
	Right()
	Report()
	Invalid(line string)
}

// Parser is constructed with a source of input commands in textual form
// and a destination CommandReceiver to send parsed commands to. Note that
// the Parser's job is _solely_ to parse commands - it provides no
// validation that the command "makes sense" given the state of the receiver.
//
// Parser is whitespace- and case-insensitive. Empty lines are silently ignored.
type Parser struct {
	scanner  *bufio.Scanner
	receiver CommandReceiver
}

func NewParser(input io.Reader, receiver CommandReceiver) *Parser {
	scanner := bufio.NewScanner(input)

	// this is the default behaviour, but lets
	// make it explicit for people new to go
	scanner.Split(bufio.ScanLines)

	return &Parser{
		scanner:  scanner,
		receiver: receiver,
	}
}

// Next scans the input source for the next available command, advances
// its internal pointer to the next line and invokes the appropriate
// method on the destination CommandReceiver. Next returns true if there
// are still commands to be returned, or false on EOF.
func (p *Parser) Next() bool {
	if p.scanner.Scan() {
		line := p.scanner.Text() // next line of input, sans newlines

		if m := placeRegex.FindStringSubmatch(line); len(m) > 0 {
			x, _ := strconv.Atoi(m[1])
			y, _ := strconv.Atoi(m[2])
			orientation := orientationFromString(m[3])
			position := position{X: x, Y: y, Orientation: orientation}
			p.receiver.Place(position)
		} else if moveRegex.MatchString(line) {
			p.receiver.Move()
		} else if leftRegex.MatchString(line) {
			p.receiver.Left()
		} else if rightRegex.MatchString(line) {
			p.receiver.Right()
		} else if reportRegex.MatchString(line) {
			p.receiver.Report()
		} else if len(strings.TrimSpace(line)) > 0 { // ignore empty lines
			p.receiver.Invalid(line)
		}

		return true
	}

	return false
}
