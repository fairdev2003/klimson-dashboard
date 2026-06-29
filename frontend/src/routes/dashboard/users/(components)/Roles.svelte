<script lang="ts"></script>

<p>Roles</p>

<div>
	<p>BaseModel</p>
	<p>
		Name string `json:"name"` Permissions []Permission `json:"permissions" gorm:"foreignKey:RoleID"`
	</p>
</div>
