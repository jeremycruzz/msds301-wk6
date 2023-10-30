# msds301-wk6

### Setup (Windows 11 + Git Bash)
- **NOTE**: Executable files are included if you don't want to deal with the system dependencies
- Clone repo with `git clone git@github.com:jeremycruzz/msds301-wk6.git`
- Install GoLearn System Dependencies [instructions here](https://github.com/gonum/blas#installation)
    - Make sure you have `g++` and `MYSS2`[instructions here](https://code.visualstudio.com/docs/cpp/config-mingw)
    - Install OpenBLAS in any directory
        - `git clone https://github.com/xianyi/OpenBLAS`
        - `cd OpenBLAS`
        - `make`
    - Link to library
        - `CGO_LDFLAGS="-L/path/to/OpenBLAS -lopenblas" go install github.com/gonum/blas/cgo`

### Building executables
- Run `go build -o sync.exe ./cmd/boston_linear_regression`
- Run `go build -o async.exe ./cmd/boston_linear_regression_async`


### Running Go executable
- Run `./sync.exe`
- Run `./async.exe`

### Results


|                | **Synchronous**   | **Asynchronous**  |
|----------------|:-----------------:|:-----------------:|
| **Model A - Min MSE** | 22.466965607975983 | 22.22454648356535  |
| **Model A - Max MSE** | 32.82293524307301  | 35.41518392090458  |
| **Model A - Avg MSE** | 26.174952309117593 | 26.256246144534398 |
| **Model B - Min MSE** | 22.11739725405122  | 21.401751274376803 |
| **Model B - Max MSE** | 38.49200879493762  | 35.060062387146615 |
| **Model B - Avg MSE** | 26.589788675963973 | 26.518287649213583 |
| **Time - Min**        | 2999500 ns         | 2076900 ns         |
| **Time - Max**        | 5716900 ns         | 3343100 ns         |
| **Time - Avg**        | 4044219 ns         | 2809534 ns         |

For the experiment, I decided to use linear regression since the median value `mv` is a continuous value. I removed two columns for each of the models. Model A omitted `rooms` and `age`, while Model B omitted `tax` and `crime`. Min, max, and average Mean Squared Error (MSE) were taken to ensure that both programs showed similar results over 100 runs. Additionally, from the MSE, we can see that omitting `rooms` and `age` produces a higher MSE than omitting `tax` and `crime`.

From the results, we can see that concurrency allowed the program to finish much faster (on average 43.95% faster) than the one without. The average time for the async program was faster than the minimum time for the sync program. The use of concurrency can allow us to quickly experiment with different features in linear regression models.

For future experimentation, it would be interesting to see how quickly we could test every combination of features and if doing so would produce any benefit.
