<html lang="en">
<head>
<title>Update student</title>
    <link rel="stylesheet" href="style.css" type="text/css"/>
</head>
<body>
<h1>Обновление студента</h1>
<form name="form" action="" method="post">
  <input type="text" name="ID" id="ID" placeholder="Enter id of student">
  <input type="text" name="fullname" id="fullname" placeholder="Enter student full name">
  <input type="text" name="group" id="group" placeholder="Enter student's group">
  <input type="submit" name="update" id="update" value="update">

</form>
<?php
 if(isset($_POST['ID']))
 {

  $mysqli = new mysqli("db", "user", "password", "appDB");
  $result = $mysqli->query("UPDATE `student` set `fullname`='{$_POST['fullname']}', `group`='{$_POST['group']}' where `ID` = {$_POST['ID']};");
}
?>

</body>
</html>