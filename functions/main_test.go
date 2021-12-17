package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoolfunc1(t *testing.T) {
	assert.Equal(t, true, boolfunc1(2))
	assert.Equal(t, false, boolfunc1(3))
}

func TestBoolfunc2(t *testing.T) {
	assert.Equal(t, true, boolfunc2(false, true))
	assert.Equal(t, true, boolfunc2(true, true))
	assert.Equal(t, true, boolfunc2(false, true))
	assert.Equal(t, false, boolfunc2(false, false))
}

func TestBoollfunc4(t *testing.T) {
	assert.Equal(t, true, boollfunc4(1, 0, 1))
	assert.Equal(t, true, boollfunc4(0, 1, 1))
	assert.Equal(t, true, boollfunc4(0, 0, 1))
	assert.Equal(t, true, boollfunc4(0, 0, 0))
	assert.Equal(t, false, boollfunc4(1, 0, 0))
	assert.Equal(t, false, boollfunc4(0, 1, 0))
}

func TestFrombooltostrlfunc(t *testing.T) {
	assert.Equal(t, "true", frombooltostrlfunc(1, 0, 1))
	assert.Equal(t, "true", frombooltostrlfunc(0, 1, 1))
	assert.Equal(t, "true", frombooltostrlfunc(0, 0, 1))
	assert.Equal(t, "true", frombooltostrlfunc(0, 0, 0))
	assert.Equal(t, "false", frombooltostrlfunc(1, 0, 0))
	assert.Equal(t, "false", frombooltostrlfunc(0, 1, 0))
}

func TestComparefunc1(t *testing.T) {
	assert.Equal(t, false, comparefunc1(-5))
	assert.Equal(t, true, comparefunc1(5))
}

func TestComparefunc2(t *testing.T) {
	assert.Equal(t, "Great", comparefunc2(85))
	assert.Equal(t, "Good", comparefunc2(45))
	assert.Equal(t, "Not bad", comparefunc2(25))
	assert.Equal(t, "Bad", comparefunc2(15))
}

func TestComparefunc3(t *testing.T) {
	assert.Equal(t, "The nearest plamet is Mercury", comparefunc3(70000_001))
	assert.Equal(t, "The nearest plamet is Venus", comparefunc3(120000001))
	assert.Equal(t, "The nearest plamet is Earth", comparefunc3(180000001))
	assert.Equal(t, "The nearest plamet is Mars", comparefunc3(500000001))
	assert.Equal(t, "The nearest plamet is Jupiter", comparefunc3(1100000001))
	assert.Equal(t, "The nearest plamet is Saturn", comparefunc3(2200000001))
	assert.Equal(t, "The nearest plamet is Uranus", comparefunc3(3500000001))
	assert.Equal(t, "The nearest plamet is Neptune", comparefunc3(43000000001))
	assert.Equal(t, "It is not a solar system", comparefunc3(80000000001))
	assert.Equal(t, "", comparefunc3(-1))
}

func TestComparefunc4(t *testing.T) {
	assert.Equal(t, "The first gear", comparefunc4(10, 2000))
	assert.Equal(t, "The second gear", comparefunc4(25, 2000))
	assert.Equal(t, "The third gear", comparefunc4(40, 2000))
	assert.Equal(t, "The fourth gear", comparefunc4(55, 2000))
	assert.Equal(t, "The fifth gear", comparefunc4(70, 2000))
	assert.Equal(t, "You need to shift into a lower gear", comparefunc4(10, 1000))
	assert.Equal(t, "You need to shift into a higher gear", comparefunc4(10, 3500))
}

func TestSignalStrenght(t *testing.T) {
	assert.Equal(t, 88, signalStrenght(600))
}

func TestDownloadTime(t *testing.T) {
	assert.Equal(t, 80.0, downloadTime(1, 100))
}
