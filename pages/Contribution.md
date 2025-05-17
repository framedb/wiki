- > FrameDB can be considered public domain, and any use of copyrighted materials is fair use here because there's no commercial activity.
  >
  >See also: [[Content Policy]]
- # Contributing to #FrameDB
	- Contribution is helping #FrameDB by adding, updating, verifying, correcting our wiki
	- Since FrameDB is open-source, anyone can contribute
	- But since FrameDB founders are mostly programmers, we encourage the use of Git and GitHub infrastructure to help with contribution integration
		- FrameDB is built on top of [logsex](https://github.com/soyart/logsex), so we'd need a Logseq graph and a GitHub repository to deploy SPA from
		- We can just use Logseq to edit/contribute new changes
		- New changes (Markdown texts) are pushed to the wiki repository, and GitHub Actions will build a new SPA and deploy it to GitHub Pages
- # Contribution steps
	- [Get Logseq](https://logseq.com)
		- Go and download Logseq for yourself
		- On macOS it should be as simple as:
		  ```shell
		  brew install logseq
		  ```
	- Go to [FrameDB GitHub](https://github.com/FrameDB/wiki), and fork or clone the repository, which will be your local repository
		- For example, you are forking the wiki to `github.com/YourUsername/wiki`
		- Then you must clone this repository down to your local machine:
		  ```shell
		  git clone https://github.com/YourUsername/wiki
		  ```
		- Once cloned, go inside and [setup Git commit hooks according to README](https://github.com/FrameDB/wiki/blob/master/README.md)
	- Open Logseq, and have it load graph from your local repository
	- [Contribute your shi](((67f4edfa-cb3d-4f92-8ef1-6e2e7a5fe1ef)))
	- Once done, commit it the changes in your local repository, and push it to the fork.
	- [Then open PR to FrameDB on GitHub](https://docs.github.com/en/pull-requests/collaborating-with-pull-requests/proposing-changes-to-your-work-with-pull-requests/creating-a-pull-request-from-a-fork)
- # Contribute
  id:: 67f4edfa-cb3d-4f92-8ef1-6e2e7a5fe1ef
	- You are free to create any new Logseq page according to our rough guidelines
	- For example, if you're adding a new page to a frame model XYZ from manufacturer ABC, then you would:
		- Create a page for manufacturer `ABC`
			- Some history if you know
			- Some other facts
			- Some references, like brochures or documents
		- Create a page for frame `ABC XYZ`
			- Some features of the frame
			- Some examples of the frame
		- Make sure to link them together