<html lang="en">
<head>
<title>Read students</title>
    <link rel="stylesheet" href="style.css" type="text/css"/>
</head>
<body>
<h1>Read all values</h1>
<table>
    <tr><th>Id</th><th>FULLNAME</th><th>GROUP</th></tr>
<?php
$mysqli = new mysqli("db", "user", "password", "appDB");
$result = $mysqli->query("SELECT * FROM student");
foreach ($result as $row){
    echo "<tr><td>{$row['ID']}</td><td>{$row['fullname']}</td><td>{$row['group']}</td></tr>";
}
?>
</table>
</body>
</html>