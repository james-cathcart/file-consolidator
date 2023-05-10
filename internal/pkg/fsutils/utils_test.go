package fsutils

import (
	"testing"
)

func TestHashFile(t *testing.T) {

	type data struct {
		file1Path string
		file2Path string
		same      bool
	}

	tests := []struct {
		description string
		input       data
		expect      data
	}{
		{
			description: `With two files containing the same content`,
			input: data{
				file1Path: `../../../data/foo.txt`,
				file2Path: `../../../data/bar.txt`,
			},
			expect: data{
				same: true,
			},
		},
		{
			description: `With two files containing different content`,
			input: data{
				file1Path: `../../../data/foo.txt`,
				file2Path: `../../../data/zed.txt`,
			},
			expect: data{
				same: false,
			},
		},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {

			fooHash, err := HashFile(test.input.file1Path)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			barHash, err := HashFile(test.input.file2Path)
			if err != nil {
				t.Errorf("error: %v", err)
			}

			if test.expect.same {
				t.Log(`Hashes should match`)
				{
					if fooHash != barHash {
						t.Errorf("\tFAIL -> hashes do not match")
					} else {
						t.Logf("\tSuccess")
					}
				}
			} else {
				t.Log(`Hashes should not match`)
				{
					if fooHash != barHash {
						t.Logf("\tSuccess")
					} else {
						t.Errorf("\tFAIL -> hashes match")
					}
				}
			}
		})
	}
}
