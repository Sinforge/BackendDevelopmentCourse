from django.urls import path
from . import views

urlpatterns = [
    path('', views.greet),
    path('<str:name>/', views.greet),
]