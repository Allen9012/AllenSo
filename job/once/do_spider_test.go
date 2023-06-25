package once

import (
	"testing"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/6/23
  @desc:
  @modified by:
**/

func TestFetchInitPostList(t *testing.T) {
	if err := FetchInitPostList(); err != nil {
		t.Errorf("FetchInitPostList err: %v", err)
	}
}

func TestGetPicture(t *testing.T) {
}
