package task3

import (
	"awesomeProject/module3/task1/wordz"
	"awesomeProject/utility"
	"fmt"
	"github.com/IAKabanov/utils"
	utilsV2 "github.com/IAKabanov/utils/v2"
	utilsV3 "github.com/IAKabanov/utils/v3"
)

func Task() {
	utility.AnnounceNewTask(3, 3)

	var isExistV3 = utilsV3.InSliceInt([]int{1, 2, 3, 4, 5}, 5)
	if isExistV3 {
		fmt.Println("Using utilsV3.InSliceInt and find value")
		return
	}

	var isExistV2 = utilsV2.InSlice(wordz.Words, "Two")
	if isExistV2 {
		fmt.Println("Using utilsV2.InSlice and find value ")
		return
	}

	var isExist = utils.Contains(wordz.Words, "Two")
	if isExist {
		fmt.Println("Slice Words contain finding value")
		return
	}

	isExist = utils.ContainsInt([]int{1, 2, 3, 4, 5}, 5)
	if isExist {
		fmt.Println("Slice Int contain finding value")
		return
	}
}
