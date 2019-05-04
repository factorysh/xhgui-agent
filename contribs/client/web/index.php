<html>

<head>
    <title>Demo</title>
</head>

<body>

    <?php

    $name = $_GET['name'] ?: 'world';
    echo "<h1>Hello $name</h1>";

    function fibo($n)
    {
        return (($n < 2) ? 1 : fibo($n - 1) + fibo($n - 2));
    }

    $n = $_GET['n'] ?: 10;
    $f = fibo($n);
    echo "Fibonacci($n) = $f";
    ?>
</body>

</html>