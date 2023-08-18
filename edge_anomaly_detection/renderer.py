import os
from jinja2 import Template

# Load the YAML template from the file
with open('detect_service.yaml') as f:
    template = Template(f.read())

# Define the variables for rendering
service_name = 'my-service'
target_port = 8080

# Render the template with the variables
rendered_yaml = template.render(SERVICE_NAME=service_name, TARGET_PORT=target_port)

# Save the rendered YAML to a new file
with open('my-service.yaml', 'w') as f:
    f.write(rendered_yaml)