# This is an auto-generated Django model module.
# You'll have to do the following manually to clean this up:
#   * Rearrange models' order
#   * Make sure each model has one field with primary_key=True
#   * Make sure each ForeignKey and OneToOneField has `on_delete` set to the desired behavior
#   * Remove `managed = False` lines if you wish to allow Django to create, modify, and delete the table
# Feel free to rename the models, but don't rename db_table values or field names.
from django.db import models


class Assignments(models.Model):
    course = models.ForeignKey('Courses', models.DO_NOTHING, blank=True, null=True)
    title = models.CharField(max_length=255)
    description = models.TextField(blank=True, null=True)
    deadline = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'assignments'


class Courses(models.Model):
    title = models.CharField(max_length=255)
    description = models.TextField(blank=True, null=True)
    teacher = models.ForeignKey('Users', models.DO_NOTHING, blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'courses'


class Materials(models.Model):
    course = models.ForeignKey(Courses, models.DO_NOTHING, blank=True, null=True)
    title = models.CharField(max_length=255)
    file_path = models.TextField()

    class Meta:
        managed = False
        db_table = 'materials'


class StudentUploads(models.Model):
    assignment = models.ForeignKey(Assignments, models.DO_NOTHING, blank=True, null=True)
    student = models.ForeignKey('Users', models.DO_NOTHING, blank=True, null=True)
    file_path = models.TextField()
    submitted_at = models.DateTimeField(blank=True, null=True)

    class Meta:
        managed = False
        db_table = 'student_uploads'


class Users(models.Model):
    name = models.CharField(max_length=100, blank=True, null=True)
    email = models.CharField(unique=True, max_length=100)
    password = models.CharField(max_length=255)
    role = models.CharField(max_length=20)

    class Meta:
        managed = False
        db_table = 'users'
