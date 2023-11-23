from django.urls import path
from . import views

urlpatterns = [
    path('', views.greet),
    path("analytics/diagram/", views.makeDiagram),
    path('<str:name>/', views.greet)
]