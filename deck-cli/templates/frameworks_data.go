package templates

var Frameworks_template = map[string]interface{}{
	"templates": map[string]interface{}{
		"flask": map[string]interface{}{
			"framework":        "flask",
			"base_image":       "python:3.11-slim",
			"work_dir":         "/app",
			"requirements_file": "requirements.txt",
			"port":             "8000",
			"run_command":      "python -m flask run --host=0.0.0.0 --port={{.Port}}",
		},
		"express": map[string]interface{}{
			"framework":        "express",
			"base_image":       "node:18-slim",
			"work_dir":         "/usr/src/app/",
			"requirements_file": "package.json",
			"port":             "3000",
			"run_command":      "npm start --host=0.0.0.0",
		},
		"react": map[string]interface{}{
			"framework":        "react",
			"base_image":       "node:18-slim",
			"work_dir":         "/usr/src/app/",
			"requirements_file": "package.json",
			"port":             "3000",
			"run_command":      "npm run start --host=0.0.0.0",
		},
		"angular": map[string]interface{}{
			"framework":        "angular",
			"base_image":       "node:18-slim",
			"work_dir":         "/usr/src/app/",
			"requirements_file": "package.json",
			"port":             "4200",
			"run_command":      "npx ng serve --host 0.0.0.0",
		},
		"django": map[string]interface{}{
			"framework":        "django",
			"base_image":       "python:3.11-slim",
			"work_dir":         "/app",
			"requirements_file": "requirements.txt",
			"port":             "8000",
			"run_command":      "python manage.py runserver",
		},
		"fastapi": map[string]interface{}{
			"framework":        "fastapi",
			"base_image":       "python:3.11-slim",
			"work_dir":         "/app",
			"requirements_file": "requirements.txt",
			"port":             "8000",
			"run_command":      "uvicorn app:app --host=0.0.0.0",
		},
		"rails": map[string]interface{}{
			"framework":        "rails",
			"base_image":       "ruby:3.2-slim",
			"work_dir":         "/app",
			"requirements_file": "Gemfile",
			"run_command":      "rails server -b 0.0.0.0 -p",
		},
		"spring_boot": map[string]interface{}{
			"framework":        "spring_boot",
			"base_image":       "openjdk:17-jdk-slim",
			"work_dir":         "/app",
			"requirements_file": "pom.xml",
			"run_command":      "java -jar target/app.jar",
		},
		"laravel": map[string]interface{}{
			"framework":        "laravel",
			"base_image":       "php:8.2-apache",
			"work_dir":         "/var/www/html",
			"requirements_file": "composer.json",
			"run_command":      "apache2-foreground",
		},
		"nextjs": map[string]interface{}{
			"framework":        "nextjs",
			"base_image":       "node:18-slim",
			"work_dir":         "/usr/src/app/",
			"requirements_file": "package.json",
			"run_command":      "npm run start --host=0.0.0.0",
		},
	},
}


var DockerIgnoreTemplate = map[string][]string{
	"react":   {"node_modules/", "build/", ".env", "Dockerfile", ".git", "public/", "logs/", "coverage/"},
	"angular": {"node_modules/", "dist/", ".env", "Dockerfile", ".git", "logs/", "coverage/"},
	"nodejs":  {"node_modules/", "npm-debug.log", "logs/", ".env", "Dockerfile", ".git", "coverage/"},
	"flask":   {"__pycache__/", ".env", "venv/", "Dockerfile", ".git", "logs/", "instance/", "test.db"},
	"fastapi": {"__pycache__/", ".env", "venv/", "Dockerfile", ".git", "logs/", "instance/", "test.db"},
	"django":  {"__pycache__/", ".env", "venv/", "Dockerfile", ".git", "logs/", "db.sqlite3", "media/", "staticfiles/", "instance/", "*.pyc"},
}
