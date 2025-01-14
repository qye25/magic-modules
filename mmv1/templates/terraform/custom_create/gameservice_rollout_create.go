// Store the ID now
id, err := ReplaceVars(d, config, "projects/{{project}}/locations/global/gameServerDeployments/{{deployment_id}}/rollout")
if err != nil {
	return fmt.Errorf("Error constructing id: %s", err)
}
d.SetId(id)

log.Printf("[DEBUG] Creating GameServerDeploymentRollout %q: ", d.Id())

err = resourceGameServicesGameServerDeploymentRolloutUpdate(d, meta)
if err != nil {
	d.SetId("")
	return fmt.Errorf("Error trying to create GameServerDeploymentRollout: %s", err)
}

return nil
