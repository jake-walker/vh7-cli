#!/usr/bin/env node
import { Command } from 'commander';
import signale from 'signale';
import prompts from 'prompts';
import clipboardy from 'clipboardy';
import version from 'project-version';
import vh7 from './vh7';

const program = new Command();

program.version(version);

program.command('shorten <url>')
  .description('Shorten a given URL')
  .action(async (url) => {
    signale.await(`Shortening ${url}...`);

    let shortlink;
    try {
      shortlink = await vh7.createShorten(url);
    } catch (error) {
      signale.error('There was a problem shortening!', error.message);
      return;
    }

    signale.success('Shortened to', `https://vh7.uk/${shortlink.link}`);
  });

// program.command('upload <path>')
//     .description('Upload a given file')
//     .action(async (file) => {

//     });

program.command('paste')
  .description('Save the clipboard contents')
  .option('-f --force', 'do not ask')
  .action(async () => {
    let content = '';

    try {
      content = await clipboardy.read();
    } catch {
      signale.error('Failed to read clipboard!');
      return;
    }

    if (content.trim() === '') {
      signale.error('There is nothing on your clipboard!');
      return;
    }

    const preview = content.length > 50 ? `${content.substr(0, 50).trim()}...` : content.trim();
    signale.info('Your clipboard contains:', preview);

    const languages = (await vh7.getLanguages()).map((lang) => ({
      title: lang.name,
      value: lang.id,
    }));

    const response = await prompts([
      {
        type: 'confirm',
        name: 'confirm',
        message: 'Are you sure you would like to save your clipboard contents?',
        initial: false,
      },
      {
        type: (prev) => (prev === true ? 'autocomplete' : null),
        name: 'langauge',
        message: 'What language is your paste?',
        choices: languages,
        initial: languages.findIndex((lang: any) => lang.value === 'plaintext'),
      },
    ]);

    if (!response.confirm) {
      return;
    }

    signale.await('Creating paste...');
    let shortlink;
    try {
      shortlink = await vh7.createPaste(response.langauge, content);
    } catch (error) {
      signale.error('There was a problem creating the paste!', error.message);
      return;
    }
    signale.success('Pasted to', `https://vh7.uk/${shortlink.link}`);
  });

program.parse(process.argv);
