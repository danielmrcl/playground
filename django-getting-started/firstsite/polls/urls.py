from polls import views
from django.urls import path

urlpatterns = [
    path('', views.index),
    path('<int:question_id>/', views.details),
    path('<int:question_id>/results/', views.results),
    path('<int:question_id>/votes/', views.votes)
]
