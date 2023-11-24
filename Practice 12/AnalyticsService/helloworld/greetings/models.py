from django.db import models

# Create your models here.
from django.conf import settings

class Analytics(models.Model):
    id = models.AutoField(primary_key=True)
    product_type = models.CharField(max_length=20)
    sum = models.IntegerField()
    count = models.IntegerField()
    image = models.BinaryField()

