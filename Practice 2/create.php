<html lang="en">
<head>
<title>Creating page</title>
    <link rel="stylesheet" href="style.css" type="text/css"/>
</head>
<body>
<h1>Создание студента</h1>
<form name="form" action="create.php" method="post">
  <input type="text" name="fullname" id="fullname" placeholder="Enter student full name">
  <input type="text" name="group" id="group" placeholder="Enter student's group">
  <input type="submit" name="save" id="save" values="save">
</form>
<?php
 if(isset($_POST['fullname']))
 {
  $mysqli = new mysqli("db", "user", "password", "appDB");
  $result = $mysqli->query("INSERT INTO `student` (`fullname`, `group`) values ('{$_POST['fullname']}', '{$_POST['group']}');");
 }

?>

</body>
</html>