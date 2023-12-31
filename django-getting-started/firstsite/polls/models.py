from datetime import timedelta
from django.db import models
from django.utils import timezone

class Question(models.Model):
    question_text = models.CharField(max_length=200)
    pub_date = models.DateTimeField(auto_now_add=True)

    def was_published_recently(self):
        return self.pub_date >= timezone.now() - timedelta(days=1)

    def __str__(self):
        return f"{self.id} - {self.question_text} - {self.pub_date}"

class Choice(models.Model):
    question = models.ForeignKey(to=Question, on_delete=models.SET_NULL, null=True)
    choice_text = models.CharField(max_length=200)
    votes = models.IntegerField(default=0)

    def __str__(self):
        return f"{self.id} - {self.question.id} - {self.choice_text} - {self.votes}"
