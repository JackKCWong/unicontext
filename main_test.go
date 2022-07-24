package unicontext_test

import (
	"context"
	"testing"
	"time"

	"github.com/JackKCWong/unicontext"
)


func TestTimeout(t *testing.T) {
	uc := unicontext.WithTimeOut(context.Background(), 1 * time.Second)

	select {
	case <- uc.Done():
		t.Fatal("unexpected timeout before 1s")
	default:
		t.Log("continue")
	}

	time.Sleep(2 * time.Second)
	select {
	case <- uc.Done():
		t.Log("2 sec passed")
	default:
		t.Fatal("unexpected timeout before 2s")
	}

	uc.Reset()
	select {
	case <- uc.Done():
		t.Fatal("unexpected timeout")
	default:
		t.Log("reset success")
	}

	uc.Cancel()
	select {
	case <- uc.Done():
		t.Log("cancel success")
	default:
		t.Fatal("cancel failed")
	}

}
