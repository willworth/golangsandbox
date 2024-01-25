Go was created out of the need to get work done. It’s not the latest trend in programming language theory, but it is a way to solve real-world problems.

It draws concepts from imperative languages with static typing. It’s fast to compile and fast to execute, it adds easy-to-understand concurrency because multi-core CPUs are now common, and it’s used successfully in large codebases (~100 million loc at Google, Inc.).

Go comes with a good standard library and a sizeable community.

```
<span></span><span>// Single line comment</span>
<span>/* Multi-</span>
<span> line comment */</span>

 <span>/* A build tag is a line comment starting with // +build</span>
<span>  and can be executed by go build -tags="foo bar" command.</span>
<span>  Build tags are placed before the package clause near or at the top of the file</span>
<span>  followed by a blank line or other line comments. */</span>
<span>// +build prod, dev, test</span>

<span>// A package clause starts every source file.</span>
<span>// main is a special name declaring an executable rather than a library.</span>
<span>package</span> <span>main</span>

<span>// Import declaration declares library packages referenced in this file.</span>
<span>import</span> <span>(</span>
    <span>"fmt"</span>       <span>// A package in the Go standard library.</span>
    <span>"io/ioutil"</span> <span>// Implements some I/O utility functions.</span>
    <span>m</span> <span>"math"</span>    <span>// Math library with local alias m.</span>
    <span>"net/http"</span>  <span>// Yes, a web server!</span>
    <span>"os"</span>        <span>// OS functions like working with the file system</span>
    <span>"strconv"</span>   <span>// String conversions.</span>
<span>)</span>

<span>// A function definition. Main is special. It is the entry point for the</span>
<span>// executable program. Love it or hate it, Go uses brace brackets.</span>
<span>func</span> <span>main</span><span>()</span> <span>{</span>
    <span>// Println outputs a line to stdout.</span>
    <span>// It comes from the package fmt.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"Hello world!"</span><span>)</span>

    <span>// Call another function within this package.</span>
    <span>beyondHello</span><span>()</span>
<span>}</span>

<span>// Functions have parameters in parentheses.</span>
<span>// If there are no parameters, empty parentheses are still required.</span>
<span>func</span> <span>beyondHello</span><span>()</span> <span>{</span>
    <span>var</span> <span>x</span> <span>int</span> <span>// Variable declaration. Variables must be declared before use.</span>
    <span>x</span> <span>=</span> <span>3</span>     <span>// Variable assignment.</span>
    <span>// "Short" declarations use := to infer the type, declare, and assign.</span>
    <span>y</span> <span>:=</span> <span>4</span>
    <span>sum</span><span>,</span> <span>prod</span> <span>:=</span> <span>learnMultiple</span><span>(</span><span>x</span><span>,</span> <span>y</span><span>)</span>        <span>// Function returns two values.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"sum:"</span><span>,</span> <span>sum</span><span>,</span> <span>"prod:"</span><span>,</span> <span>prod</span><span>)</span> <span>// Simple output.</span>
    <span>learnTypes</span><span>()</span>                            <span>// &lt; y minutes, learn more!</span>
<span>}</span>

<span>/* &lt;- multiline comment</span>
<span>Functions can have parameters and (multiple!) return values.</span>
<span>Here `x`, `y` are the arguments and `sum`, `prod` is the signature (what's returned).</span>
<span>Note that `x` and `sum` receive the type `int`.</span>
<span>*/</span>
<span>func</span> <span>learnMultiple</span><span>(</span><span>x</span><span>,</span> <span>y</span> <span>int</span><span>)</span> <span>(</span><span>sum</span><span>,</span> <span>prod</span> <span>int</span><span>)</span> <span>{</span>
    <span>return</span> <span>x</span> <span>+</span> <span>y</span><span>,</span> <span>x</span> <span>*</span> <span>y</span> <span>// Return two values.</span>
<span>}</span>

<span>// Some built-in types and literals.</span>
<span>func</span> <span>learnTypes</span><span>()</span> <span>{</span>
    <span>// Short declaration usually gives you what you want.</span>
    <span>str</span> <span>:=</span> <span>"Learn Go!"</span> <span>// string type.</span>

    <span>s2</span> <span>:=</span> <span>`A "raw" string literal</span>
<span>can include line breaks.`</span> <span>// Same string type.</span>

    <span>// Non-ASCII literal. Go source is UTF-8.</span>
    <span>g</span> <span>:=</span> <span>'Σ'</span> <span>// rune type, an alias for int32, holds a unicode code point.</span>

    <span>f</span> <span>:=</span> <span>3.14195</span> <span>// float64, an IEEE-754 64-bit floating point number.</span>
    <span>c</span> <span>:=</span> <span>3</span> <span>+</span> <span>4i</span>  <span>// complex128, represented internally with two float64's.</span>

    <span>// var syntax with initializers.</span>
    <span>var</span> <span>u</span> <span>uint</span> <span>=</span> <span>7</span> <span>// Unsigned, but implementation dependent size as with int.</span>
    <span>var</span> <span>pi</span> <span>float32</span> <span>=</span> <span>22.</span> <span>/</span> <span>7</span>

    <span>// Conversion syntax with a short declaration.</span>
    <span>n</span> <span>:=</span> <span>byte</span><span>(</span><span>'\n'</span><span>)</span> <span>// byte is an alias for uint8.</span>

    <span>// Arrays have size fixed at compile time.</span>
    <span>var</span> <span>a4</span> <span>[</span><span>4</span><span>]</span><span>int</span>           <span>// An array of 4 ints, initialized to all 0.</span>
    <span>a5</span> <span>:=</span> <span>[</span><span>...</span><span>]</span><span>int</span><span>{</span><span>3</span><span>,</span> <span>1</span><span>,</span> <span>5</span><span>,</span> <span>10</span><span>,</span> <span>100</span><span>}</span> <span>// An array initialized with a fixed size of five</span>
    <span>// elements, with values 3, 1, 5, 10, and 100.</span>

    <span>// Arrays have value semantics.</span>
    <span>a4_cpy</span> <span>:=</span> <span>a4</span>            <span>// a4_cpy is a copy of a4, two separate instances.</span>
    <span>a4_cpy</span><span>[</span><span>0</span><span>]</span> <span>=</span> <span>25</span>          <span>// Only a4_cpy is changed, a4 stays the same.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>a4_cpy</span><span>[</span><span>0</span><span>]</span> <span>==</span> <span>a4</span><span>[</span><span>0</span><span>])</span> <span>// false</span>

    <span>// Slices have dynamic size. Arrays and slices each have advantages</span>
    <span>// but use cases for slices are much more common.</span>
    <span>s3</span> <span>:=</span> <span>[]</span><span>int</span><span>{</span><span>4</span><span>,</span> <span>5</span><span>,</span> <span>9</span><span>}</span>    <span>// Compare to a5. No ellipsis here.</span>
    <span>s4</span> <span>:=</span> <span>make</span><span>([]</span><span>int</span><span>,</span> <span>4</span><span>)</span>    <span>// Allocates slice of 4 ints, initialized to all 0.</span>
    <span>var</span> <span>d2</span> <span>[][]</span><span>float64</span>      <span>// Declaration only, nothing allocated here.</span>
    <span>bs</span> <span>:=</span> <span>[]</span><span>byte</span><span>(</span><span>"a slice"</span><span>)</span> <span>// Type conversion syntax.</span>

    <span>// Slices (as well as maps and channels) have reference semantics.</span>
    <span>s3_cpy</span> <span>:=</span> <span>s3</span>            <span>// Both variables point to the same instance.</span>
    <span>s3_cpy</span><span>[</span><span>0</span><span>]</span> <span>=</span> <span>0</span>           <span>// Which means both are updated.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>s3_cpy</span><span>[</span><span>0</span><span>]</span> <span>==</span> <span>s3</span><span>[</span><span>0</span><span>])</span> <span>// true </span>

    <span>// Because they are dynamic, slices can be appended to on-demand.</span>
    <span>// To append elements to a slice, the built-in append() function is used.</span>
    <span>// First argument is a slice to which we are appending. Commonly,</span>
    <span>// the array variable is updated in place, as in example below.</span>
    <span>s</span> <span>:=</span> <span>[]</span><span>int</span><span>{</span><span>1</span><span>,</span> <span>2</span><span>,</span> <span>3</span><span>}</span>     <span>// Result is a slice of length 3.</span>
    <span>s</span> <span>=</span> <span>append</span><span>(</span><span>s</span><span>,</span> <span>4</span><span>,</span> <span>5</span><span>,</span> <span>6</span><span>)</span>  <span>// Added 3 elements. Slice now has length of 6.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>s</span><span>)</span> <span>// Updated slice is now [1 2 3 4 5 6]</span>

    <span>// To append another slice, instead of list of atomic elements we can</span>
    <span>// pass a reference to a slice or a slice literal like this, with a</span>
    <span>// trailing ellipsis, meaning take a slice and unpack its elements,</span>
    <span>// appending them to slice s.</span>
    <span>s</span> <span>=</span> <span>append</span><span>(</span><span>s</span><span>,</span> <span>[]</span><span>int</span><span>{</span><span>7</span><span>,</span> <span>8</span><span>,</span> <span>9</span><span>}</span><span>...</span><span>)</span> <span>// Second argument is a slice literal.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>s</span><span>)</span>  <span>// Updated slice is now [1 2 3 4 5 6 7 8 9]</span>

    <span>p</span><span>,</span> <span>q</span> <span>:=</span> <span>learnMemory</span><span>()</span> <span>// Declares p, q to be type pointer to int.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>*</span><span>p</span><span>,</span> <span>*</span><span>q</span><span>)</span>   <span>// * follows a pointer. This prints two ints.</span>

    <span>// Maps are a dynamically growable associative array type, like the</span>
    <span>// hash or dictionary types of some other languages.</span>
    <span>m</span> <span>:=</span> <span>map</span><span>[</span><span>string</span><span>]</span><span>int</span><span>{</span><span>"three"</span><span>:</span> <span>3</span><span>,</span> <span>"four"</span><span>:</span> <span>4</span><span>}</span>
    <span>m</span><span>[</span><span>"one"</span><span>]</span> <span>=</span> <span>1</span>

    <span>// Unused variables are an error in Go.</span>
    <span>// The underscore lets you "use" a variable but discard its value.</span>
    <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span><span>,</span> <span>_</span> <span>=</span> <span>str</span><span>,</span> <span>s2</span><span>,</span> <span>g</span><span>,</span> <span>f</span><span>,</span> <span>u</span><span>,</span> <span>pi</span><span>,</span> <span>n</span><span>,</span> <span>a5</span><span>,</span> <span>s4</span><span>,</span> <span>bs</span>
    <span>// Usually you use it to ignore one of the return values of a function</span>
    <span>// For example, in a quick and dirty script you might ignore the</span>
    <span>// error value returned from os.Create, and expect that the file</span>
    <span>// will always be created.</span>
    <span>file</span><span>,</span> <span>_</span> <span>:=</span> <span>os</span><span>.</span><span>Create</span><span>(</span><span>"output.txt"</span><span>)</span>
    <span>fmt</span><span>.</span><span>Fprint</span><span>(</span><span>file</span><span>,</span> <span>"This is how you write to a file, by the way"</span><span>)</span>
    <span>file</span><span>.</span><span>Close</span><span>()</span>

    <span>// Output of course counts as using a variable.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>s</span><span>,</span> <span>c</span><span>,</span> <span>a4</span><span>,</span> <span>s3</span><span>,</span> <span>d2</span><span>,</span> <span>m</span><span>)</span>

    <span>learnFlowControl</span><span>()</span> <span>// Back in the flow.</span>
<span>}</span>

<span>// It is possible, unlike in many other languages for functions in go</span>
<span>// to have named return values.</span>
<span>// Assigning a name to the type being returned in the function declaration line</span>
<span>// allows us to easily return from multiple points in a function as well as to</span>
<span>// only use the return keyword, without anything further.</span>
<span>func</span> <span>learnNamedReturns</span><span>(</span><span>x</span><span>,</span> <span>y</span> <span>int</span><span>)</span> <span>(</span><span>z</span> <span>int</span><span>)</span> <span>{</span>
    <span>z</span> <span>=</span> <span>x</span> <span>*</span> <span>y</span>
    <span>return</span> <span>// z is implicit here, because we named it earlier.</span>
<span>}</span>

<span>// Go is fully garbage collected. It has pointers but no pointer arithmetic.</span>
<span>// You can make a mistake with a nil pointer, but not by incrementing a pointer.</span>
<span>// Unlike in C/Cpp taking and returning an address of a local variable is also safe. </span>
<span>func</span> <span>learnMemory</span><span>()</span> <span>(</span><span>p</span><span>,</span> <span>q</span> <span>*</span><span>int</span><span>)</span> <span>{</span>
    <span>// Named return values p and q have type pointer to int.</span>
    <span>p</span> <span>=</span> <span>new</span><span>(</span><span>int</span><span>)</span> <span>// Built-in function new allocates memory.</span>
    <span>// The allocated int slice is initialized to 0, p is no longer nil.</span>
    <span>s</span> <span>:=</span> <span>make</span><span>([]</span><span>int</span><span>,</span> <span>20</span><span>)</span> <span>// Allocate 20 ints as a single block of memory.</span>
    <span>s</span><span>[</span><span>3</span><span>]</span> <span>=</span> <span>7</span>             <span>// Assign one of them.</span>
    <span>r</span> <span>:=</span> <span>-</span><span>2</span>              <span>// Declare another local variable.</span>
    <span>return</span> <span>&amp;</span><span>s</span><span>[</span><span>3</span><span>],</span> <span>&amp;</span><span>r</span>     <span>// &amp; takes the address of an object.</span>
<span>}</span>

<span>// Use the aliased math library (see imports, above) </span>
<span>func</span> <span>expensiveComputation</span><span>()</span> <span>float64</span> <span>{</span>
    <span>return</span> <span>m</span><span>.</span><span>Exp</span><span>(</span><span>10</span><span>)</span>
<span>}</span>

<span>func</span> <span>learnFlowControl</span><span>()</span> <span>{</span>
    <span>// If statements require brace brackets, and do not require parentheses.</span>
    <span>if</span> <span>true</span> <span>{</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"told ya"</span><span>)</span>
    <span>}</span>
    <span>// Formatting is standardized by the command line command "go fmt".</span>
    <span>if</span> <span>false</span> <span>{</span>
        <span>// Pout.</span>
    <span>}</span> <span>else</span> <span>{</span>
        <span>// Gloat.</span>
    <span>}</span>
    <span>// Use switch in preference to chained if statements.</span>
    <span>x</span> <span>:=</span> <span>42.0</span>
    <span>switch</span> <span>x</span> <span>{</span>
    <span>case</span> <span>0</span><span>:</span>
    <span>case</span> <span>1</span><span>,</span> <span>2</span><span>:</span> <span>// Can have multiple matches on one case</span>
    <span>case</span> <span>42</span><span>:</span>
        <span>// Cases don't "fall through".</span>
        <span>/*</span>
<span>        There is a `fallthrough` keyword however, see:</span>
<span>          https://github.com/golang/go/wiki/Switch#fall-through</span>
<span>        */</span>
    <span>case</span> <span>43</span><span>:</span>
        <span>// Unreached.</span>
    <span>default</span><span>:</span>
        <span>// Default case is optional.</span>
    <span>}</span>

    <span>// Type switch allows switching on the type of something instead of value</span>
    <span>var</span> <span>data</span> <span>interface</span><span>{}</span>
    <span>data</span> <span>=</span> <span>""</span>
    <span>switch</span> <span>c</span> <span>:=</span> <span>data</span><span>.(</span><span>type</span><span>)</span> <span>{</span>
    <span>case</span> <span>string</span><span>:</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>c</span><span>,</span> <span>"is a string"</span><span>)</span>
    <span>case</span> <span>int64</span><span>:</span>
        <span>fmt</span><span>.</span><span>Printf</span><span>(</span><span>"%d is an int64\n"</span><span>,</span> <span>c</span><span>)</span>
    <span>default</span><span>:</span>
        <span>// all other cases</span>
    <span>}</span>

    <span>// Like if, for doesn't use parens either.</span>
    <span>// Variables declared in for and if are local to their scope.</span>
    <span>for</span> <span>x</span> <span>:=</span> <span>0</span><span>;</span> <span>x</span> <span>&lt;</span> <span>3</span><span>;</span> <span>x</span><span>++</span> <span>{</span> <span>// ++ is a statement.</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"iteration"</span><span>,</span> <span>x</span><span>)</span>
    <span>}</span>
    <span>// x == 42 here.</span>

    <span>// For is the only loop statement in Go, but it has alternate forms.</span>
    <span>for</span> <span>{</span> <span>// Infinite loop.</span>
        <span>break</span>    <span>// Just kidding.</span>
        <span>continue</span> <span>// Unreached.</span>
    <span>}</span>

    <span>// You can use range to iterate over an array, a slice, a string, a map, or a channel.</span>
    <span>// range returns one (channel) or two values (array, slice, string and map).</span>
    <span>for</span> <span>key</span><span>,</span> <span>value</span> <span>:=</span> <span>range</span> <span>map</span><span>[</span><span>string</span><span>]</span><span>int</span><span>{</span><span>"one"</span><span>:</span> <span>1</span><span>,</span> <span>"two"</span><span>:</span> <span>2</span><span>,</span> <span>"three"</span><span>:</span> <span>3</span><span>}</span> <span>{</span>
        <span>// for each pair in the map, print key and value</span>
        <span>fmt</span><span>.</span><span>Printf</span><span>(</span><span>"key=%s, value=%d\n"</span><span>,</span> <span>key</span><span>,</span> <span>value</span><span>)</span>
    <span>}</span>
    <span>// If you only need the value, use the underscore as the key</span>
    <span>for</span> <span>_</span><span>,</span> <span>name</span> <span>:=</span> <span>range</span> <span>[]</span><span>string</span><span>{</span><span>"Bob"</span><span>,</span> <span>"Bill"</span><span>,</span> <span>"Joe"</span><span>}</span> <span>{</span>
        <span>fmt</span><span>.</span><span>Printf</span><span>(</span><span>"Hello, %s\n"</span><span>,</span> <span>name</span><span>)</span>
    <span>}</span>

    <span>// As with for, := in an if statement means to declare and assign</span>
    <span>// y first, then test y &gt; x.</span>
    <span>if</span> <span>y</span> <span>:=</span> <span>expensiveComputation</span><span>();</span> <span>y</span> <span>&gt;</span> <span>x</span> <span>{</span>
        <span>x</span> <span>=</span> <span>y</span>
    <span>}</span>
    <span>// Function literals are closures.</span>
    <span>xBig</span> <span>:=</span> <span>func</span><span>()</span> <span>bool</span> <span>{</span>
        <span>return</span> <span>x</span> <span>&gt;</span> <span>10000</span> <span>// References x declared above switch statement.</span>
    <span>}</span>
    <span>x</span> <span>=</span> <span>99999</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"xBig:"</span><span>,</span> <span>xBig</span><span>())</span> <span>// true</span>
    <span>x</span> <span>=</span> <span>1.3e3</span>                    <span>// This makes x == 1300</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"xBig:"</span><span>,</span> <span>xBig</span><span>())</span> <span>// false now.</span>

    <span>// What's more is function literals may be defined and called inline,</span>
    <span>// acting as an argument to function, as long as:</span>
    <span>// a) function literal is called immediately (),</span>
    <span>// b) result type matches expected type of argument.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"Add + double two numbers: "</span><span>,</span>
        <span>func</span><span>(</span><span>a</span><span>,</span> <span>b</span> <span>int</span><span>)</span> <span>int</span> <span>{</span>
            <span>return</span> <span>(</span><span>a</span> <span>+</span> <span>b</span><span>)</span> <span>*</span> <span>2</span>
        <span>}(</span><span>10</span><span>,</span> <span>2</span><span>))</span> <span>// Called with args 10 and 2</span>
    <span>// =&gt; Add + double two numbers: 24</span>

    <span>// When you need it, you'll love it.</span>
    <span>goto</span> <span>love</span>
<span>love</span><span>:</span>

    <span>learnFunctionFactory</span><span>()</span> <span>// func returning func is fun(3)(3)</span>
    <span>learnDefer</span><span>()</span>      <span>// A quick detour to an important keyword.</span>
    <span>learnInterfaces</span><span>()</span> <span>// Good stuff coming up!</span>
<span>}</span>

<span>func</span> <span>learnFunctionFactory</span><span>()</span> <span>{</span>
    <span>// Next two are equivalent, with second being more practical</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>sentenceFactory</span><span>(</span><span>"summer"</span><span>)(</span><span>"A beautiful"</span><span>,</span> <span>"day!"</span><span>))</span>

    <span>d</span> <span>:=</span> <span>sentenceFactory</span><span>(</span><span>"summer"</span><span>)</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>d</span><span>(</span><span>"A beautiful"</span><span>,</span> <span>"day!"</span><span>))</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>d</span><span>(</span><span>"A lazy"</span><span>,</span> <span>"afternoon!"</span><span>))</span>
<span>}</span>

<span>// Decorators are common in other languages. Same can be done in Go</span>
<span>// with function literals that accept arguments.</span>
<span>func</span> <span>sentenceFactory</span><span>(</span><span>mystring</span> <span>string</span><span>)</span> <span>func</span><span>(</span><span>before</span><span>,</span> <span>after</span> <span>string</span><span>)</span> <span>string</span> <span>{</span>
    <span>return</span> <span>func</span><span>(</span><span>before</span><span>,</span> <span>after</span> <span>string</span><span>)</span> <span>string</span> <span>{</span>
        <span>return</span> <span>fmt</span><span>.</span><span>Sprintf</span><span>(</span><span>"%s %s %s"</span><span>,</span> <span>before</span><span>,</span> <span>mystring</span><span>,</span> <span>after</span><span>)</span> <span>// new string</span>
    <span>}</span>
<span>}</span>

<span>func</span> <span>learnDefer</span><span>()</span> <span>(</span><span>ok</span> <span>bool</span><span>)</span> <span>{</span>
    <span>// A defer statement pushes a function call onto a list. The list of saved</span>
    <span>// calls is executed AFTER the surrounding function returns.</span>
    <span>defer</span> <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"deferred statements execute in reverse (LIFO) order."</span><span>)</span>
    <span>defer</span> <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"\nThis line is being printed first because"</span><span>)</span>
    <span>// Defer is commonly used to close a file, so the function closing the</span>
    <span>// file stays close to the function opening the file.</span>
    <span>return</span> <span>true</span>
<span>}</span>

<span>// Define Stringer as an interface type with one method, String.</span>
<span>type</span> <span>Stringer</span> <span>interface</span> <span>{</span>
    <span>String</span><span>()</span> <span>string</span>
<span>}</span>

<span>// Define pair as a struct with two fields, ints named x and y.</span>
<span>type</span> <span>pair</span> <span>struct</span> <span>{</span>
    <span>x</span><span>,</span> <span>y</span> <span>int</span>
<span>}</span>

<span>// Define a method on type pair. Pair now implements Stringer because Pair has defined all the methods in the interface.</span>
<span>func</span> <span>(</span><span>p</span> <span>pair</span><span>)</span> <span>String</span><span>()</span> <span>string</span> <span>{</span> <span>// p is called the "receiver"</span>
    <span>// Sprintf is another public function in package fmt.</span>
    <span>// Dot syntax references fields of p.</span>
    <span>return</span> <span>fmt</span><span>.</span><span>Sprintf</span><span>(</span><span>"(%d, %d)"</span><span>,</span> <span>p</span><span>.</span><span>x</span><span>,</span> <span>p</span><span>.</span><span>y</span><span>)</span>
<span>}</span>

<span>func</span> <span>learnInterfaces</span><span>()</span> <span>{</span>
    <span>// Brace syntax is a "struct literal". It evaluates to an initialized</span>
    <span>// struct. The := syntax declares and initializes p to this struct.</span>
    <span>p</span> <span>:=</span> <span>pair</span><span>{</span><span>3</span><span>,</span> <span>4</span><span>}</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>p</span><span>.</span><span>String</span><span>())</span> <span>// Call String method of p, of type pair.</span>
    <span>var</span> <span>i</span> <span>Stringer</span>          <span>// Declare i of interface type Stringer.</span>
    <span>i</span> <span>=</span> <span>p</span>                   <span>// Valid because pair implements Stringer</span>
    <span>// Call String method of i, of type Stringer. Output same as above.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>i</span><span>.</span><span>String</span><span>())</span>

    <span>// Functions in the fmt package call the String method to ask an object</span>
    <span>// for a printable representation of itself.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>p</span><span>)</span> <span>// Output same as above. Println calls String method.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>i</span><span>)</span> <span>// Output same as above.</span>

    <span>learnVariadicParams</span><span>(</span><span>"great"</span><span>,</span> <span>"learning"</span><span>,</span> <span>"here!"</span><span>)</span>
<span>}</span>

<span>// Functions can have variadic parameters.</span>
<span>func</span> <span>learnVariadicParams</span><span>(</span><span>myStrings</span> <span>...</span><span>interface</span><span>{})</span> <span>{</span>
    <span>// Iterate each value of the variadic.</span>
    <span>// The underbar here is ignoring the index argument of the array.</span>
    <span>for</span> <span>_</span><span>,</span> <span>param</span> <span>:=</span> <span>range</span> <span>myStrings</span> <span>{</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"param:"</span><span>,</span> <span>param</span><span>)</span>
    <span>}</span>

    <span>// Pass variadic value as a variadic parameter.</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"params:"</span><span>,</span> <span>fmt</span><span>.</span><span>Sprintln</span><span>(</span><span>myStrings</span><span>...</span><span>))</span>

    <span>learnErrorHandling</span><span>()</span>
<span>}</span>

<span>func</span> <span>learnErrorHandling</span><span>()</span> <span>{</span>
    <span>// ", ok" idiom used to tell if something worked or not.</span>
    <span>m</span> <span>:=</span> <span>map</span><span>[</span><span>int</span><span>]</span><span>string</span><span>{</span><span>3</span><span>:</span> <span>"three"</span><span>,</span> <span>4</span><span>:</span> <span>"four"</span><span>}</span>
    <span>if</span> <span>x</span><span>,</span> <span>ok</span> <span>:=</span> <span>m</span><span>[</span><span>1</span><span>];</span> <span>!</span><span>ok</span> <span>{</span> <span>// ok will be false because 1 is not in the map.</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"no one there"</span><span>)</span>
    <span>}</span> <span>else</span> <span>{</span>
        <span>fmt</span><span>.</span><span>Print</span><span>(</span><span>x</span><span>)</span> <span>// x would be the value, if it were in the map.</span>
    <span>}</span>
    <span>// An error value communicates not just "ok" but more about the problem.</span>
    <span>if</span> <span>_</span><span>,</span> <span>err</span> <span>:=</span> <span>strconv</span><span>.</span><span>Atoi</span><span>(</span><span>"non-int"</span><span>);</span> <span>err</span> <span>!=</span> <span>nil</span> <span>{</span> <span>// _ discards value</span>
        <span>// prints 'strconv.ParseInt: parsing "non-int": invalid syntax'</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>err</span><span>)</span>
    <span>}</span>
    <span>// We'll revisit interfaces a little later. Meanwhile,</span>
    <span>learnConcurrency</span><span>()</span>
<span>}</span>

<span>// c is a channel, a concurrency-safe communication object.</span>
<span>func</span> <span>inc</span><span>(</span><span>i</span> <span>int</span><span>,</span> <span>c</span> <span>chan</span> <span>int</span><span>)</span> <span>{</span>
    <span>c</span> <span>&lt;-</span> <span>i</span> <span>+</span> <span>1</span> <span>// &lt;- is the "send" operator when a channel appears on the left.</span>
<span>}</span>

<span>// We'll use inc to increment some numbers concurrently.</span>
<span>func</span> <span>learnConcurrency</span><span>()</span> <span>{</span>
    <span>// Same make function used earlier to make a slice. Make allocates and</span>
    <span>// initializes slices, maps, and channels.</span>
    <span>c</span> <span>:=</span> <span>make</span><span>(</span><span>chan</span> <span>int</span><span>)</span>
    <span>// Start three concurrent goroutines. Numbers will be incremented</span>
    <span>// concurrently, perhaps in parallel if the machine is capable and</span>
    <span>// properly configured. All three send to the same channel.</span>
    <span>go</span> <span>inc</span><span>(</span><span>0</span><span>,</span> <span>c</span><span>)</span> <span>// go is a statement that starts a new goroutine.</span>
    <span>go</span> <span>inc</span><span>(</span><span>10</span><span>,</span> <span>c</span><span>)</span>
    <span>go</span> <span>inc</span><span>(</span><span>-</span><span>805</span><span>,</span> <span>c</span><span>)</span>
    <span>// Read three results from the channel and print them out.</span>
    <span>// There is no telling in what order the results will arrive!</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>&lt;-</span><span>c</span><span>,</span> <span>&lt;-</span><span>c</span><span>,</span> <span>&lt;-</span><span>c</span><span>)</span> <span>// channel on right, &lt;- is "receive" operator.</span>

    <span>cs</span> <span>:=</span> <span>make</span><span>(</span><span>chan</span> <span>string</span><span>)</span>       <span>// Another channel, this one handles strings.</span>
    <span>ccs</span> <span>:=</span> <span>make</span><span>(</span><span>chan</span> <span>chan</span> <span>string</span><span>)</span> <span>// A channel of string channels.</span>
    <span>go</span> <span>func</span><span>()</span> <span>{</span> <span>c</span> <span>&lt;-</span> <span>84</span> <span>}()</span>       <span>// Start a new goroutine just to send a value.</span>
    <span>go</span> <span>func</span><span>()</span> <span>{</span> <span>cs</span> <span>&lt;-</span> <span>"wordy"</span> <span>}()</span> <span>// Again, for cs this time.</span>
    <span>// Select has syntax like a switch statement but each case involves</span>
    <span>// a channel operation. It selects a case at random out of the cases</span>
    <span>// that are ready to communicate.</span>
    <span>select</span> <span>{</span>
    <span>case</span> <span>i</span> <span>:=</span> <span>&lt;-</span><span>c</span><span>:</span> <span>// The value received can be assigned to a variable,</span>
        <span>fmt</span><span>.</span><span>Printf</span><span>(</span><span>"it's a %T"</span><span>,</span> <span>i</span><span>)</span>
    <span>case</span> <span>&lt;-</span><span>cs</span><span>:</span> <span>// or the value received can be discarded.</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"it's a string"</span><span>)</span>
    <span>case</span> <span>&lt;-</span><span>ccs</span><span>:</span> <span>// Empty channel, not ready for communication.</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>"didn't happen."</span><span>)</span>
    <span>}</span>
    <span>// At this point a value was taken from either c or cs. One of the two</span>
    <span>// goroutines started above has completed, the other will remain blocked.</span>

    <span>learnWebProgramming</span><span>()</span> <span>// Go does it. You want to do it too.</span>
<span>}</span>

<span>// A single function from package http starts a web server.</span>
<span>func</span> <span>learnWebProgramming</span><span>()</span> <span>{</span>

    <span>// First parameter of ListenAndServe is TCP address to listen to.</span>
    <span>// Second parameter is an interface, specifically http.Handler.</span>
    <span>go</span> <span>func</span><span>()</span> <span>{</span>
        <span>err</span> <span>:=</span> <span>http</span><span>.</span><span>ListenAndServe</span><span>(</span><span>":8080"</span><span>,</span> <span>pair</span><span>{})</span>
        <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>err</span><span>)</span> <span>// don't ignore errors</span>
    <span>}()</span>

    <span>requestServer</span><span>()</span>
<span>}</span>

<span>// Make pair an http.Handler by implementing its only method, ServeHTTP.</span>
<span>func</span> <span>(</span><span>p</span> <span>pair</span><span>)</span> <span>ServeHTTP</span><span>(</span><span>w</span> <span>http</span><span>.</span><span>ResponseWriter</span><span>,</span> <span>r</span> <span>*</span><span>http</span><span>.</span><span>Request</span><span>)</span> <span>{</span>
    <span>// Serve data with a method of http.ResponseWriter.</span>
    <span>w</span><span>.</span><span>Write</span><span>([]</span><span>byte</span><span>(</span><span>"You learned Go in Y minutes!"</span><span>))</span>
<span>}</span>

<span>func</span> <span>requestServer</span><span>()</span> <span>{</span>
    <span>resp</span><span>,</span> <span>err</span> <span>:=</span> <span>http</span><span>.</span><span>Get</span><span>(</span><span>"http://localhost:8080"</span><span>)</span>
    <span>fmt</span><span>.</span><span>Println</span><span>(</span><span>err</span><span>)</span>
    <span>defer</span> <span>resp</span><span>.</span><span>Body</span><span>.</span><span>Close</span><span>()</span>
    <span>body</span><span>,</span> <span>err</span> <span>:=</span> <span>ioutil</span><span>.</span><span>ReadAll</span><span>(</span><span>resp</span><span>.</span><span>Body</span><span>)</span>
    <span>fmt</span><span>.</span><span>Printf</span><span>(</span><span>"\nWebserver said: `%s`"</span><span>,</span> <span>string</span><span>(</span><span>body</span><span>))</span>
<span>}</span>
```

The root of all things Go is the [official Go web site](http://golang.org/). There you can follow the tutorial, play interactively, and read lots. Aside from a tour, [the docs](https://golang.org/doc/) contain information on how to write clean and effective Go code, package and command docs, and release history.

The [Go language specification](https://golang.org/ref/spec) itself is highly recommended. It’s easy to read and amazingly short (as language definitions go these days.)

You can play around with the code on [Go playground](https://play.golang.org/p/tnWMjr16Mm). Try to change it and run it from your browser! Note that you can use [https://play.golang.org](https://play.golang.org/) as a [REPL](https://en.wikipedia.org/wiki/Read-eval-print_loop) to test things and code in your browser, without even installing Go.

On the reading list for students of Go is the [source code to the standard library](http://golang.org/src/pkg/). Comprehensively documented, it demonstrates the best of readable and understandable Go, Go style, and Go idioms. Or you can click on a function name in [the documentation](http://golang.org/pkg/) and the source code comes up!

There are many excellent conference talks and video tutorials on Go available on YouTube, and here are three playlists of the very best, tailored for beginners, intermediate, and advanced Gophers respectively:

Go Mobile adds support for mobile platforms (Android and iOS). You can write all-Go native mobile apps or write a library that contains bindings from a Go package, which can be invoked via Java (Android) and Objective-C (iOS). Check out the [Go Mobile page](https://github.com/golang/go/wiki/Mobile) for more information.

Got a suggestion? A correction, perhaps? [Open an Issue](https://github.com/adambard/learnxinyminutes-docs/issues/new) on the Github Repo, or make a [pull request](https://github.com/adambard/learnxinyminutes-docs/edit/master/go.html.markdown) yourself!