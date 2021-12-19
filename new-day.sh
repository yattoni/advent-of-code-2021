if [ "$#" -ne 1 ]; then
    echo "Illegal number of parameters"
	exit
fi

mkdir $1

cd $1

touch README.md

touch prompt-input

touch answers.txt

touch main.go

cat > main.go << EOF
package $1

import (
    "fmt"
)

func Solve(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}
EOF

touch main_test.go

cat > main_test.go << EOF
package $1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolve(t *testing.T) {
	assert.Equal(t, "Hello, World!", Solve("World"))
}
EOF

