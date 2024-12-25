package templates


var Frameworks_template = map[string]interface{}{

	"templates" : map[string]interface{}{
		"flask": map[string]interface{}{
			"framework": "flask",
			"base_image": "python:3.11-slim",
			"work_dir": "/app",
			"requirements_file": "requirements.txt",
			"run_command": "python -m flask run --port=3001",
		},
		"express": map[string]interface{}{
			"framework": "express",
			"base_image": "node:18-slim",
			"work_dir": "/usr/src/app/",
			"requirements_file": "package.json",
			"run_command": "npm start",
		},
		"react": map[string]interface{}{
			"framework": "react",
			"base_image": "node:18-slim",
			"work_dir": "/usr/src/app/",
			"requirements_file": "package.json",
			"run_command": "npm run start",
		},
		"angular": map[string]interface{}{
			"framework": "angular",
			"base_image": "node:18-slim",
			"work_dir": "/usr/src/app/",
			"requirements_file": "package.json",
			"run_command": "ng serve --host 0.0.0.0 --port 4200",
		},
		"django": map[string]interface{}{
			"framework": "django",
			"base_image": "python:3.11-slim",
			"work_dir": "/app",
			"requirements_file": "requirements.txt",
			"run_command": "python manage.py runserver 0.0.0.0:8000",
		},
		"fastapi": map[string]interface{}{
			"framework": "fastapi",
			"base_image": "python:3.11-slim",
			"work_dir": "/app",
			"requirements_file": "requirements.txt",
			"run_command": "uvicorn app.main:app --host 0.0.0.0 --port 8000",
		},
		"rails": map[string]interface{}{
			"framework": "rails",
			"base_image": "ruby:3.2-slim",
			"work_dir": "/app",
			"requirements_file": "Gemfile",
			"run_command": "rails server -b 0.0.0.0 -p 3000",
		},
		"spring_boot": map[string]interface{}{
			"framework": "spring_boot",
			"base_image": "openjdk:17-jdk-slim",
			"work_dir": "/app",
			"requirements_file": "pom.xml",
			"run_command": "java -jar target/app.jar",
		},
		"laravel": map[string]interface{}{
			"framework": "laravel",
			"base_image": "php:8.2-apache",
			"work_dir": "/var/www/html",
			"requirements_file": "composer.json",
			"run_command": "apache2-foreground",
		},
		"nextjs": map[string]interface{}{
			"framework": "nextjs",
			"base_image": "node:18-slim",
			"work_dir": "/usr/src/app/",
			"requirements_file": "package.json",
			"run_command": "npm run start",
		},
	},
}

// Mapping of development templates to their corresponding frameworks
var DockerfileDevTemplates = map[string]string{
	"flask":   FlaskDockerfileDevTemplate,
	"django":  DjangoDockerfileDevTemplate,
	"nodejs":  NodejsDockerfileDevTemplate,
	"react":   ReactDockerfileDevTemplate,
	"golang":  GolangDockerfileDevTemplate,
	"angular": AngularDockerfileDevTemplate,
	"fastapi": FastAPIDockerfileDevTemplate,
	"express": ExpressDockerfileDevTemplate,
}