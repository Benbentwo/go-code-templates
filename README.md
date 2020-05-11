These files are used by the `Generate Function File` Command of Ben's Binary (`bb`) to generate code for a go repo.

![Some Profound Text](https://raw.githubusercontent.com/Benbentwo/bb/master/docs/AvBinaryGenCode.gif)

> NOTE: releasing this as v0 because indentation might be horked but it worked a while ago. I intend to clean it up, but the first step is to move it to a more available location.
>         The other reason for v0 is I'm trying to remove `github.com/jenkins-x/jx` as a dependency for templates. Adding Jenkins X as a dependency adds a TON of indirect modules to a project.
>         Thirdly. The reason for version controlling through file names, is so I have a migraiton plan, v0 can still continue to function while I release new versions. 

##### v0.0.0

 - Determine if these should be in a repository
     - `+` Better Source Control
     - `-` Bloatting Repo Numbers to host small number of files